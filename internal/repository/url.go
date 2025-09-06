package repository

import (
	"fmt"

	"github.com/nikitaenmi/URLShortener/internal/domain"
	"gorm.io/gorm"
)

type Url struct {
	DB *gorm.DB
}

func NewUrl(db *gorm.DB) Url {
	return Url{
		DB: db,
	}
}

func (r Url) URLFind(params domain.URLFilter) (*domain.Url, error) {
	url := domain.Url{}

	q := r.DB.Select("*")
	q = r.buildFilterByParams(q, params)

	if err := q.First(&url).Error; err != nil {
		return nil, err
	}

	return &url, nil
}

func (r Url) buildFilterByParams(q *gorm.DB, params domain.URLFilter) *gorm.DB {
	if params.Alias != "" {
		q.Where(&domain.Url{Alias: params.Alias})
	}

	return q
}

func (r Url) Create(url domain.Url) error {
	if err := r.DB.Create(&url).Error; err != nil {
		return fmt.Errorf("repo create error: %w", err)
	}

	return nil
}
