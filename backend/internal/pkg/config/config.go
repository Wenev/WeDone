package config

import "os"

type Config struct {
	POSTGRES_CONN string
	PORT string
}

func Load() *Config {
	return &Config{
		POSTGRES_CONN: getEnv("POSTGRES_CONN", ""),
		PORT: getEnv("PORT", "3000"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}