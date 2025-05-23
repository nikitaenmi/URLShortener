package repository

import (
	"errors"

	"github.com/nikitaenmi/URLShortener/internal/domain"
	"gorm.io/gorm"
)

type UrlDB struct {
	DB *gorm.DB
}

func (r *UrlDB) URLFind(alias string) (*domain.Link, error) {
	var link domain.Link

	result := r.DB.Where("alias = ?", alias).First(&link)
	if result.Error != nil {
		return nil, result.Error
	}

	return &link, nil
}

func (r *UrlDB) Create(URL, alias string) error {
	link := domain.Link{
		OriginalURL: URL,
		Alias:       alias,
	}

	err := r.DB.Create(&link).Error

	if err != nil {
		return errors.New("Repo create error: " + err.Error())
	}
	return nil
}
