package repository

import (
	"context"
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

func (r Url) URLFind(ctx context.Context, params domain.URLFilter) (*domain.Url, error) {
	url := domain.Url{}

	q := r.DB.WithContext(ctx).Select("*")
	q = r.buildFilterByParams(q, params)

	if err := q.First(&url).Error; err != nil {
		return nil, err
	}
	return &url, nil
}

func (r Url) buildFilterByParams(q *gorm.DB, params domain.URLFilter) *gorm.DB {
	if params.Alias != "" {
		q = q.Where(&domain.Url{Alias: params.Alias})
	}

	if params.ID != "" {
		q = q.Where(&domain.Url{ID: params.ID})
	}
	return q
}

func (r Url) Create(ctx context.Context, url domain.Url) error {
	if err := r.DB.WithContext(ctx).Create(&url).Error; err != nil {
		return fmt.Errorf("repo create error: %w", err)
	}

	return nil
}

func (r Url) Delete(ctx context.Context, params domain.URLFilter) error {
	result := r.DB.WithContext(ctx).Where("ID = ?", params.ID).Delete(&domain.Url{})
	if result.Error != nil {
		return fmt.Errorf("failed deleting url in database: %w", result.Error)
	}
	return nil
}
