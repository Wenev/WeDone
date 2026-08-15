package config

import "os"

type Config struct {
	PostgresConn string `json:"postgres_conn"`
	Port         string `json:"port"`
}

func Load() *Config {
	return &Config{
		PostgresConn: getEnv("POSTGRES_CONN", ""),
		Port:         getEnv("PORT", "3000"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
