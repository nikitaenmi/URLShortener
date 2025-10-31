package repository

import (
	"context"
	"fmt"

	"github.com/nikitaenmi/URLShortener/internal/domain"
	"gorm.io/gorm"
)

type URL struct {
	DB *gorm.DB
}

func NewURL(db *gorm.DB) URL {
	return URL{
		DB: db,
	}
}

func (r URL) Create(ctx context.Context, url domain.URL) error {
	if err := r.DB.WithContext(ctx).Create(&url).Error; err != nil {
		return fmt.Errorf("repo create error: %w", err)
	}

	return nil
}

func (r URL) buildFilterByParams(q *gorm.DB, params domain.URLFilter) *gorm.DB {
	if params.Alias != "" {
		q = q.Where(&domain.URL{Alias: params.Alias})
	}

	if params.ID != 0 {
		q = q.Where(&domain.URL{ID: params.ID})
	}
	return q
}

func (r URL) Update(ctx context.Context, url *domain.URL) error {
	params := domain.URLFilter{
		ID: url.ID,
	}
	q := r.DB.WithContext(ctx).Model(&domain.URL{})
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

func (r URL) Delete(ctx context.Context, params domain.URLFilter) error {
	q := r.DB.WithContext(ctx)
	q = r.buildFilterByParams(q, params)

	res := q.Delete(&domain.URL{})
	if res.Error != nil {
		return fmt.Errorf("failed deleting url in database: %w", res.Error)
	}
	if res.RowsAffected == 0 {
		return domain.ErrURLNotFound
	}
	return nil
}

func (r URL) Count(ctx context.Context, params domain.URLFilter) (int64, error) {
	var count int64

	q := r.DB.WithContext(ctx).Model(&domain.URL{})
	q = r.buildFilterByParams(q, params)

	if err := q.Count(&count).Error; err != nil {
		return 0, fmt.Errorf("repo count error: %w", err)
	}

	return count, nil
}

func (r URL) FindAll(ctx context.Context, params domain.URLFilter, paginator *domain.Paginator) ([]*domain.URL, error) {
	var urls []*domain.URL
	q := r.DB.WithContext(ctx)
	q = r.buildFilterByParams(q, params)

	if params.ID != 0 {
		q = q.Limit(1)
	} else {
		q = q.Offset(paginator.GetOffset()).Limit(paginator.GetLimit())
	}

	res := q.Order("id ASC").Find(&urls)
	if res.Error != nil {
		return nil, fmt.Errorf("repo find all error: %w", res.Error)
	}

	if res.RowsAffected == 0 {
		return nil, domain.ErrURLNotFound
	}

	return urls, nil
}
