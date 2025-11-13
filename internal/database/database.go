package database

import (
	"fmt"

	"github.com/nikitaenmi/URLShortener/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg config.Database) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("open connection error: %w", err)
	}

	err = db.AutoMigrate(&URL{})
	if err != nil {
		return nil, fmt.Errorf("migrate error: %w", err)
	}

	return db, nil
}
