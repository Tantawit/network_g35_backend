package config

import (
	gosdk "github.com/2110336-2565-2/cu-freelance-library"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type RabbitMQConfig struct {
	RabbitMQ gosdk.RabbitMQConfig `mapstructure:"rabbitmq"`
}

func LoadRabbitMQConfig() (config *gosdk.RabbitMQConfig, err error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("rabbitmq")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, errors.Wrap(err, "error occurs while reading the config")
	}

	conf := RabbitMQConfig{}

	err = viper.Unmarshal(&conf)
	if err != nil {
		return nil, errors.Wrap(err, "error occurs while unmarshal the config")
	}

	return &conf.RabbitMQ, nil
}
