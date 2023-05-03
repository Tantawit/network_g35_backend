package main

import (
	"context"
	"fmt"
	"github.com/2110336-2565-2/cu-freelance-chat/src/config"
	"github.com/2110336-2565-2/cu-freelance-chat/src/constant"
	"github.com/2110336-2565-2/cu-freelance-chat/src/metric"
	"github.com/2110336-2565-2/cu-freelance-chat/src/pkg/handler"
	"github.com/2110336-2565-2/cu-freelance-chat/src/pkg/publisher"
	"github.com/2110336-2565-2/cu-freelance-chat/src/pkg/repository"
	"github.com/2110336-2565-2/cu-freelance-chat/src/pkg/router"
	"github.com/2110336-2565-2/cu-freelance-chat/src/pkg/service"
	"github.com/2110336-2565-2/cu-freelance-chat/src/pkg/strategy"
	gosdk "github.com/2110336-2565-2/cu-freelance-library"
	"github.com/2110336-2565-2/cu-freelance-library/pkg/pb"
	"github.com/google/uuid"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"sync"
	"syscall"
	"time"
)

type operation func(ctx context.Context) error

func gracefulShutdown(ctx context.Context, logger gosdk.Logger, timeout time.Duration, ops map[string]operation) <-chan struct{} {
	wait := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)

		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
		sig := <-s

		logger.Info().
			Keyword("statement", "graceful shutdown").
			Keyword("signal", sig).
			Msg("shutting down service")

		timeoutFunc := time.AfterFunc(timeout, func() {
			logger.Error(errors.New("timeout")).
				Keyword("statement", "graceful shutdown").
				Keyword("total_time", timeout).
				Msg("timeout has been elapsed, force exit")
			os.Exit(0)
		})

		defer timeoutFunc.Stop()

		var wg sync.WaitGroup

		for key, op := range ops {
			wg.Add(1)
			innerOp := op
			innerKey := key
			go func() {
				defer wg.Done()

				logger.Info().
					Keyword("statement", "graceful shutdown").
					Keyword("inner_key", innerKey).
					Msg("cleaning up")
				if err := innerOp(ctx); err != nil {
					logger.Error(err).
						Keyword("statement", "graceful shutdown").
						Keyword("inner_key", innerKey).
						Msg("%v: clean up failed")
					return
				}

				logger.Info().
					Keyword("statement", "graceful shutdown").
					Keyword("inner_key", innerKey).
					Msg("shutdown gracefully")
			}()
		}

		wg.Wait()
		close(wait)
	}()

	return wait
}

func main() {
	logger := gosdk.NewLogger("cu-freelance-chat")
	constant.ServerID = gosdk.UUIDAdr(uuid.New())

	conf, err := config.LoadAppConfig()
	if err != nil {
		logger.
			Fatal(err).
			Keyword("statement", "load config").
			Msg("failed to start service")
	}

	if conf.Sentry.DSN != "" {
		if err := gosdk.SetSentryDSN(conf.Sentry.DSN); err != nil {
			logger.
				Fatal(err).
				Keyword("statement", "set sentry dsn").
				Msg("failed to start service")
		}
	}

	if err := gosdk.SetUpTracer(conf.JaegerConfig); err != nil {
		logger.
			Fatal(err).
			Keyword("host", conf.JaegerConfig.Host).
			Keyword("environment", conf.JaegerConfig.Environment).
			Keyword("service_name", conf.JaegerConfig.ServiceName).
			Msg("failed to setup jaeger")
	}

	if err := metric.SetupMetric(); err != nil {
		logger.
			Fatal(err).
			Keyword("environment", conf.JaegerConfig.Environment).
			Msg("failed to setup metric")
	}

	//db, err := gosdk.InitCassandraConnection(conf.CassandraConfig)
	//if err != nil {
	//	logger.
	//		Fatal(err).
	//		Keyword("statement", "init cassandra database connection").
	//		Msg("failed to start service")
	//}

	chatSessionDB, err := gosdk.InitRedisConnect(conf.ChatSessionRedis)
	if err != nil {
		logger.
			Fatal(err).
			Keyword("url", conf.Service.Account).
			Msg("failed to connect to chat service")
	}

	accountConn, err := grpc.Dial(
		conf.Service.Account,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(gosdk.NewGRPUnaryClientInterceptor()),
	)
	if err != nil {
		logger.
			Fatal(err).
			Keyword("url", conf.Service.Account).
			Msg("failed to connect to chat service")
	}

	rabbitMQConn, err := gosdk.InitRabbitMQConnection(conf.RabbitMQConfig)
	if err != nil {
		logger.Fatal(err).
			Keyword("action", "init rabbitmq client").
			Keyword("v_host", conf.RabbitMQConfig.VHost).
			Keyword("host", conf.RabbitMQConfig.Host).
			Msg("failed to init rabbitmq connection")
	}

	grpcLis, err := net.Listen("tcp", fmt.Sprintf(":%v", conf.App.GRPCPort))
	if err != nil {
		logger.
			Fatal(err).
			Keyword("statement", "listen to tcp").
			Keyword("port", conf.App.GRPCPort).
			Msg("failed to start service")
	}

	grpcPanicRecoveryHandler := func(p any) (err error) {
		metric.PanicsTotal.Inc()
		logger.Error(errors.New("received panic")).
			Keyword("error", p).
			Keyword("stack", string(debug.Stack())).
			Msg("recovered from panic")
		return status.Errorf(codes.Internal, "%s", p)
	}

	chatSessionRepository := repository.NewChatSessionRepository(chatSessionDB)
	//chatPrivateRepository := repository.NewChatPrivateRepository(db)
	//chatPrivateService := service.NewChatPrivateService(chatPrivateRepository)
	//chatHandler := handler.NewWebSocketChatHandler(chatPrivateService)

	client := pb.NewAuthServiceClient(accountConn)
	authService := service.NewAuthService(client)

	chatPub, err := publisher.NewChatPublisher(rabbitMQConn)
	if err != nil {
		logger.
			Fatal(err).
			Keyword("statement", "init portfolio publisher").
			Keyword("host", conf.RabbitMQConfig.Host).
			Keyword("v_host", conf.RabbitMQConfig.VHost).
			Msg("failed to init portfolio publisher")
	}
	wsChatStrategy := strategy.NewWebsocketChatStrategy(chatPub, chatSessionRepository)
	wsChatService := service.NewWebsocketChatService(wsChatStrategy)
	wsChatHandler := handler.NewWebSocketChatHandler(authService, wsChatService, conf.App.KeepAliveInterval)

	r := router.NewFiberRouter()
	r.WebsocketChat(wsChatHandler.Listen)
	v1 := router.NewAPIv1(r, conf.App)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			gosdk.NewGRPUnaryServerInterceptor(),
			metric.ServerMetrics.UnaryServerInterceptor(),
			recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler)),
		),
	)

	metric.ServerMetrics.InitializeMetrics(grpcServer)

	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())

	reflection.Register(grpcServer)

	pMux := http.NewServeMux()
	pMux.Handle("/metrics", promhttp.Handler())

	go func() {
		logger.
			Info().
			Keyword("statement", "start the service").
			Keyword("port", conf.App.MetricPort).
			Msg("[Metric] cufreelance chat starting")

		if err := http.ListenAndServe(fmt.Sprintf(":%v", conf.App.MetricPort), pMux); err != nil {
			logger.
				Fatal(err).
				Keyword("statement", "serving metric").
				Keyword("port", conf.App.MetricPort).
				Msg("failed to expose metrics endpoint")
		}
	}()

	go func() {
		logger.
			Info().
			Keyword("statement", "start the service").
			Keyword("port", conf.App.GRPCPort).
			Msg("[gRPC] cufreelance chat starting")

		if err = grpcServer.Serve(grpcLis); err != nil {
			logger.
				Fatal(err).
				Keyword("statement", "serving grpc").
				Keyword("port", conf.App.GRPCPort).
				Msg("[gRPC] failed to start service")
		}
	}()

	go func() {
		logger.
			Info().
			Keyword("statement", "start the service").
			Keyword("port", conf.App.WebSocketPort).
			Msg("[WebSocket] starting service")

		if err = v1.Listen(fmt.Sprintf(":%v", conf.App.WebSocketPort)); err != nil {
			logger.
				Fatal(err).
				Keyword("statement", "serving grpc").
				Keyword("port", conf.App.WebSocketPort).
				Msg("[WebSocket] fail to start service")
		}
	}()

	wait := gracefulShutdown(context.Background(), logger, 2*time.Second, map[string]operation{
		//"database": func(ctx context.Context) error {
		//	db.Close()
		//	return nil
		//},
		"grpc": func(ctx context.Context) error {
			grpcServer.GracefulStop()
			return nil
		},
		"ws": func(ctx context.Context) error {
			return r.Shutdown()
		},
		"jaeger": func(ctx context.Context) error {
			return gosdk.CloseTracer()
		},
	})

	<-wait

	logger.
		Info().
		Keyword("statement", "graceful shutdown").
		Msg("service terminated")
}
