package models

import (
	"github.com/teris-io/shortid"
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

func (r *UrlDB) FinderAlias(alias string) (*Link, error) { // по короткому коду, который был сгенерированный, возращает оригинальную ссылку
	var link Link

	result := r.DB.Where("alias = ?", alias).First(&link)
	if result.Error != nil {
		return nil, result.Error
	}

	return &link, nil
}

// ShortIDGenerator структура, метод которой генерит алиас
type AliasGenerator struct{}

func (g *AliasGenerator) Generate() (string, error) {
	return shortid.Generate()
}
