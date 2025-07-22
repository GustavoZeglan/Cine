package config

import (
	"fmt"
	"sync"

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

	OracleDB struct {
		Host     string
		Name     string
		Port     int
		User     string
		Password string
	}
}

var (
	AppConfig *Config
	once      sync.Once
)

func LoadConfig() *Config {
	once.Do(func() {

		v := viper.New()

		v.SetConfigName("config")
		v.SetConfigType("yaml")
		v.AddConfigPath(".")

		v.AutomaticEnv()

		if err := v.ReadInConfig(); err != nil {
			fmt.Println("Error reading config file", err)
		}

		if err := v.Unmarshal(&AppConfig); err != nil {
			fmt.Println("Error unmarshalling config file", err)
		}

	})

	return AppConfig
}
