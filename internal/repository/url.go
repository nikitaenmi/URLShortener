package repository

import (
	"context"
	"errors"
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrURLNotFound
		}
		return nil, fmt.Errorf("repo find error: %w", err)
	}
	return &url, nil
}

func (r Url) Update(ctx context.Context, url *domain.Url) error {
	params := domain.URLFilter{
		ID: url.ID,
	}
	q := r.DB.WithContext(ctx).Model(&domain.Url{})
	q = r.buildFilterByParams(q, params)

	res := q.Update("original_url", url.OriginalURL)

	if res.Error != nil {
		return fmt.Errorf("failed to update URL in database: %w", res.Error)
	}

	if res.RowsAffected == 0 {
		return domain.ErrURLNotFound
	}

	return nil
}

func (r Url) Delete(ctx context.Context, params domain.URLFilter) error {
	res := r.DB.WithContext(ctx).Where("ID = ?", params.ID).Delete(&domain.Url{})
	if res.Error != nil {
		return fmt.Errorf("failed deleting url in database: %w", res.Error)
	}
	if res.RowsAffected == 0 {
		return domain.ErrURLNotFound
	}
	return nil
}

func (r Url) Count(ctx context.Context, params domain.URLFilter) (int64, error) {
	var count int64

	q := r.DB.WithContext(ctx).Model(&domain.Url{})
	q = r.buildFilterByParams(q, params)

	if err := q.Count(&count).Error; err != nil {
		return 0, fmt.Errorf("repo count error: %w", err)
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

	res := q.Order("id ASC").Find(&urls)
	if res.Error != nil {
		return nil, fmt.Errorf("repo find all error: %w", res.Error)
	}
	return urls, nil
}
