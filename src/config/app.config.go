package config

import (
	gosdk "github.com/2110336-2565-2/cu-freelance-library"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Service struct {
	Account string `json:"account"`
}

type App struct {
	WebSocketPort     int  `mapstructure:"http-port"`
	GRPCPort          int  `mapstructure:"grpc-port"`
	MetricPort        int  `mapstructure:"metric-port"`
	Debug             bool `mapstructure:"debug"`
	KeepAliveInterval int  `mapstructure:"keep-alive-interval"`
}

type Config struct {
	App              App     `mapstructure:"app"`
	Service          Service `mapstructure:"service"`
	ChatSessionRedis *gosdk.RedisConfig
	*gosdk.CassandraConfig
	*gosdk.JaegerConfig
	*gosdk.RabbitMQConfig
	*Sentry
}

func LoadAppConfig() (config *Config, err error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, errors.Wrap(err, "error occurs while reading the config")
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, errors.Wrap(err, "error occurs while unmarshal the config")
	}

	cassandraConf, err := LoadDatabaseConfig()
	if err != nil {
		return nil, err
	}

	redisConf, err := LoadRedisConfig()
	if err != nil {
		return nil, err
	}

	sentryConf, err := LoadSentryConfig()
	if err != nil {
		return nil, err
	}

	rabbitmqConf, err := LoadRabbitMQConfig()
	if err != nil {
		return nil, err
	}

	jaegerConf, err := LoadJaegerConfig()
	if err != nil {
		return nil, err
	}

	config.CassandraConfig = cassandraConf
	config.Sentry = sentryConf
	config.JaegerConfig = jaegerConf
	config.ChatSessionRedis = redisConf.ChatSession
	config.RabbitMQConfig = rabbitmqConf

	return
}
