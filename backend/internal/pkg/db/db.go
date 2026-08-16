package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresConnection(conn_string string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(conn_string), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}

func NewSQLiteConnection(conn_string string) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(conn_string))
}