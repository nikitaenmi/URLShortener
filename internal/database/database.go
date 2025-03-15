package database

import (
	"errors"
	"fmt"

	"github.com/nikitaenmi/URLShortener/internal/config"
	"github.com/nikitaenmi/URLShortener/internal/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg config.Database) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)), &gorm.Config{})

	if err != nil {
		return nil, errors.New("Open connection error: " + err.Error())

	}

	err = db.AutoMigrate(&models.Link{})
	if err != nil {
		return nil, errors.New("Migrate error:" + err.Error())
	}

	return db, nil
}
