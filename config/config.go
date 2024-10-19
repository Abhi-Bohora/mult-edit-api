package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host string
		Port string
		User string
		Password string
		DBName string
		SSLMode string
	}

	Server struct {
		Port string
	}
}

func LoadConfig() (*Config, error) {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("./config")

    var config Config

    if err := viper.ReadInConfig(); err != nil {
        return nil, err
    }

    if err := viper.Unmarshal(&config); err != nil {
        return nil, err
    }

    return &config, nil
}