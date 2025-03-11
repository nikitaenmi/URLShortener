package models

import (
	"gorm.io/gorm"
)

type Link struct {
	ID          int `gorm:"primarykey"`
	OriginalURL string
	Aliace      string
}

type UrlDB struct {
	DB *gorm.DB
}

func (r *UrlDB) FinderOriginalCode(alias string) (*Link, error) { // по короткому коду, который был сгенерированный, возращает оригинальную ссылку
	var link Link

	result := r.DB.Where("aliace = ?", alias).First(&link)
	if result.Error != nil {
		return nil, result.Error
	}

	return &link, nil
}
