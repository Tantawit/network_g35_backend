package config

import (
	gosdk "github.com/2110336-2565-2/cu-freelance-library"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type JaegerConfig struct {
	Jaeger *gosdk.JaegerConfig `mapstructure:"jaeger"`
}

func LoadJaegerConfig() (*gosdk.JaegerConfig, error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("jaeger")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "error occurs while reading the config")
	}

	conf := JaegerConfig{}

	if err := viper.Unmarshal(&conf); err != nil {
		return nil, errors.Wrap(err, "error occurs while unmarshal the config")
	}

	return conf.Jaeger, nil
}
