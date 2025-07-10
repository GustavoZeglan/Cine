package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	PostgresDB struct {
		Host     string
		Name     string
		Port     int
		User     string
		Password string
	}
}

var AppConfig *Config

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file", err)
	}

	var cf Config
	if err := viper.Unmarshal(&cf); err != nil {
		fmt.Println("Error unmarshalling config file", err)
	}

	AppConfig = &cf
}
