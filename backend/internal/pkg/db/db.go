package db

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PoolConfig struct {
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
	IsDev           bool
}

func NewPostgresConnection(conn_string string, cfg *PoolConfig) (*gorm.DB, error) {
	logLevel := logger.Warn
	if cfg.IsDev {
		logLevel = logger.Info
	}

	gormDB, err := gorm.Open(postgres.Open(conn_string), &gorm.Config{
		Logger:      logger.Default.LogMode(logLevel),
		PrepareStmt: false,
	})

	if err != nil {
		return nil, err
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	sqlDB.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)

	return gormDB, nil
}

func NewSQLiteConnection(conn_string string) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(conn_string))
}
