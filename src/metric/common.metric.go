package metric

import (
	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var ServerMetrics = grpcprom.NewServerMetrics(
	grpcprom.WithServerHandlingTimeHistogram(
		grpcprom.WithHistogramBuckets([]float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120}),
	),
)

var PanicsTotal = promauto.NewCounter(prometheus.CounterOpts{
	Name: "grpc_req_panics_recovered_total",
	Help: "Total number of gRPC requests recovered from internal panic.",
})

func SetupMetric() error {
	if err := prometheus.Register(ServerMetrics); err != nil {
		return err
	}

	return nil
}
