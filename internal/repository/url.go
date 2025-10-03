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

func (r Url) Create(ctx context.Context, url domain.Url) error {
	if err := r.DB.WithContext(ctx).Create(&url).Error; err != nil {
		return fmt.Errorf("repo create error: %w", err)
	}

	return nil
}

func (r Url) buildFilterByParams(q *gorm.DB, params domain.URLFilter) *gorm.DB {
	if params.Alias != "" {
		q = q.Where(&domain.Url{Alias: params.Alias})
	}

	if params.ID != 0 {
		q = q.Where(&domain.Url{ID: params.ID})
	}
	return q
}

func (r Url) FindById(ctx context.Context, params domain.URLFilter) (*domain.Url, error) {
	url := domain.Url{}
	q := r.DB.WithContext(ctx)
	q = r.buildFilterByParams(q, params)

	if err := q.First(&url).Error; err != nil {
		return nil, err
	}
	return &url, nil
}

func (r Url) Update(ctx context.Context, url *domain.Url) error {
	params := domain.URLFilter{
		ID: url.ID,
	}
	q := r.DB.WithContext(ctx).Model(&domain.Url{})
	q = r.buildFilterByParams(q, params)

	result := q.Update("original_url", url.OriginalURL)

	if result.Error != nil {
		return fmt.Errorf("failed to update URL in database: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("URL with id %d not found", url.ID)
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

func (r Url) Count(ctx context.Context, params domain.URLFilter) (int64, error) {
	var count int64

	q := r.DB.WithContext(ctx).Model(&domain.Url{})
	q = r.buildFilterByParams(q, params)

	if err := q.Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (r Url) FindAll(ctx context.Context, params domain.URLFilter, paginator *domain.Paginator) ([]*domain.Url, error) {
	var urls []*domain.Url
	q := r.DB.WithContext(ctx)
	q = r.buildFilterByParams(q, params)

	if paginator != nil {
		offset := (paginator.Page - 1) * paginator.Limit
		q = q.Offset(offset).Limit(paginator.Limit)
	}

	result := q.Order("id ASC").Find(&urls)
	return urls, result.Error
}

func (r Url) List(ctx context.Context, params domain.Paginator) ([]*domain.Url, domain.Paginator, error) {
	filter := domain.URLFilter{}

	total, err := r.Count(ctx, filter)
	if err != nil {
		return nil, params, err
	}
	params.Total = total

	urls, err := r.FindAll(ctx, filter, &params)
	if err != nil {
		return nil, params, err
	}

	return urls, params, nil
}
