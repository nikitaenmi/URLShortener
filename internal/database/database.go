package database

import (
	"fmt"
	"os"
	"urlShortener/data"
	mod "urlShortener/internal/database/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Migration() *gorm.DB {
	db, err := gorm.Open(postgres.Open(data.Dsn), &gorm.Config{})
	db.AutoMigrate(&mod.Link{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return db
}
