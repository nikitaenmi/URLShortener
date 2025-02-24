package database

import (
	"fmt"
	"os"

	"github.com/nikitaenmi/URLShortener/data"
	"github.com/nikitaenmi/URLShortener/internal/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Migration() *gorm.DB {
	db, err := gorm.Open(postgres.Open(data.Dsn), &gorm.Config{})
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
