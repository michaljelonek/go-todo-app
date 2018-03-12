package config

import (
	"os"
)

type Config struct {
	DBConfig *DBConfig
}

type DBConfig struct {
	Dialect  string
	Name     string
	Username string
	Password string
	Host     string
	Port     string
}

func getenv(key, fallback string) string {
	if value := os.Getenv(key); len(value) != 0 {
		return value
	}
	return fallback
}

func GetConf() *Config {
	return &Config{
		DBConfig: &DBConfig{
			Dialect:  getenv("TODO_DB_DIALECT", "postgres"),
			Name:     getenv("TODO_DB_NAME", "tododb"),
			Username: getenv("TODO_DB_USERNAME", "tododb"),
			Password: getenv("TODO_DB_PASSWORD", "tododb"),
			Host:     getenv("TODO_DB_HOST", "localhost"),
			Port:     getenv("TODO_DB_PORT", "5432"),
		},
	}
}
