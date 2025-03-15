package models

import (
	"errors"

	"gorm.io/gorm"
)

type Link struct {
	ID          int `gorm:"primarykey"`
	OriginalURL string
	Alias       string
}

type UrlDB struct {
	DB *gorm.DB
}

func (r *UrlDB) URLFind(alias string) (*Link, error) {
	var link Link

	result := r.DB.Where("alias = ?", alias).First(&link)
	if result.Error != nil {
		return nil, result.Error
	}

	return &link, nil
}

func (r *UrlDB) Create(URL, alias string) error {
	link := Link{
		OriginalURL: URL,
		Alias:       alias,
	}

	err := r.DB.Create(&link).Error

	if err != nil {
		return errors.New("Repo create error: " + err.Error())
	}
	return nil
}
