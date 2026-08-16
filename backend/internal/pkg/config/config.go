package config

import "os"

type Config struct {
	PostgresConn string `json:"postgres_conn"`
	Port         string `json:"port"`
	RedisAddr string
	RedisPass string
	RedisDB int
}

func Load() *Config {
	return &Config{
		PostgresConn: getEnv("POSTGRES_CONN", ""),
		Port:         getEnv("PORT", "3000"),
		RedisAddr: getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPass: getEnv("REDIS_PASS", ""),
		RedisDB: 0,
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
