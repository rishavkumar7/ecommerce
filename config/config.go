package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	MONGO_URI string
	PORT      string
}

func SetConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Config{
		MONGO_URI: viper.GetString("MONGO_URI"),
		PORT:      viper.GetString("PORT"),
	}, nil
}
