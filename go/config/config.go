package config

import (
	"github.com/spf13/viper"
)

type Config struct {
    Port     string
    Database struct {
        Driver string
        DSN    string
    }
}


func Load() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}
