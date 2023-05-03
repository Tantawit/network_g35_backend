package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Sentry struct {
	DSN string `mapstructure:"dsn"`
}

type SentryConfig struct {
	Sentry Sentry `mapstructure:"sentry"`
}

func LoadSentryConfig() (config *Sentry, err error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("sentry")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, errors.Wrap(err, "error occurs while reading the config")
	}

	conf := SentryConfig{}

	err = viper.Unmarshal(&conf)
	if err != nil {
		return nil, errors.Wrap(err, "error occurs while unmarshal the config")
	}

	return &conf.Sentry, err
}
