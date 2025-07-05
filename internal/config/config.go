package config

import "os"

type DbConn struct {
	Host     string
	DBName   string
	Port     string
	User     string
	Password string
}

func GetDatabaseEnvironments() DbConn {
	return DbConn{
		Host:     os.Getenv("POSTGRES_HOST"),
		DBName:   os.Getenv("POSTGRES_DB"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	}
}
