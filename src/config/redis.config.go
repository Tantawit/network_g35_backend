package config

import (
	gosdk "github.com/2110336-2565-2/cu-freelance-library"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type RedisConfig struct {
	ChatSession *gosdk.RedisConfig `mapstructure:"redis-chat-session"`
}

func LoadRedisConfig() (*RedisConfig, error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("redis")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "error occurs while reading the config")
	}

	conf := RedisConfig{}

	if err := viper.Unmarshal(&conf); err != nil {
		return nil, errors.Wrap(err, "error occurs while unmarshal the config")
	}

	return &conf, nil
}
