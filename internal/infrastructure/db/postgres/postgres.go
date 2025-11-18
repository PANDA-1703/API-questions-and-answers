package postgres

import (
	"fmt"

	"github.com/PANDA-1703/API-questions-and-answers/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(cfgPsql *config.PostgresConfig) (*gorm.DB, error) {
	dsn := cfgPsql.PgSource()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("gorm.Open: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("db.DB(): %w", err)
	}

	if err = sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("db.Ping: %w", err)
	}

	return db, nil
}
