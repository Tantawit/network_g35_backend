package config

import (
	gosdk "github.com/2110336-2565-2/cu-freelance-library"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Cassandra struct {
	CassandraConfig *gosdk.CassandraConfig `mapstructure:"cassandra"`
}

func LoadDatabaseConfig() (*gosdk.CassandraConfig, error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("database")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, errors.Wrap(err, "error occurs while reading the config")
	}

	conf := Cassandra{}

	err = viper.Unmarshal(&conf)
	if err != nil {
		return nil, errors.Wrap(err, "error occurs while unmarshal the config")
	}

	return conf.CassandraConfig, nil
}
