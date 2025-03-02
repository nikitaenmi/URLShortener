package database

import (
	"fmt"
	"os"

	"github.com/nikitaenmi/URLShortener/internal/config"
	"github.com/nikitaenmi/URLShortener/internal/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Migration(cfg config.Database) *gorm.DB {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.AutoMigrate(&models.Link{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return db
}
