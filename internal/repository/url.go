package repository

import (
	"context"
	"fmt"

	"github.com/nikitaenmi/URLShortener/internal/domain"
	"gorm.io/gorm"
)

const (
	idColumn          = "id"
	originalURLColumn = "original_url"
)

type URL struct {
	db *gorm.DB
}

func NewURL(db *gorm.DB) URL {
	return URL{
		db: db,
	}
}

func (r URL) Create(ctx context.Context, url domain.URL) error {
	if err := r.db.WithContext(ctx).Create(&url).Error; err != nil {
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
	q := r.db.WithContext(ctx).Model(&domain.URL{})
	q = r.buildFilterByParams(q, params)

	res := q.Update(originalURLColumn, url.OriginalURL)

	if res.Error != nil {
		return fmt.Errorf("failed to update URL in database: %w", res.Error)
	}

	if res.RowsAffected == 0 {
		return domain.ErrURLNotFound
	}

	return nil
}

func (r URL) Delete(ctx context.Context, params domain.URLFilter) error {
	q := r.db.WithContext(ctx)
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

	q := r.db.WithContext(ctx).Model(&domain.URL{})
	q = r.buildFilterByParams(q, params)

	if err := q.Count(&count).Error; err != nil {
		return 0, fmt.Errorf("repo count error: %w", err)
	}

	return count, nil
}

func (r URL) FindAll(ctx context.Context, params domain.URLFilter, paginator *domain.Paginator) ([]*domain.URL, error) {
	var urls []*domain.URL
	q := r.db.WithContext(ctx)
	q = r.buildFilterByParams(q, params)

	if paginator == nil {
		q = q.Limit(1)
	} else {
		offset := paginator.GetOffset()
		q = q.Offset(offset)
		limit := paginator.GetLimit()
		q = q.Limit(limit)
	}

	res := q.Order(idColumn + " " + "ASC").Find(&urls)
	if res.Error != nil {
		return nil, fmt.Errorf("repo find all error: %w", res.Error)
	}

	if res.RowsAffected == 0 {
		return nil, domain.ErrURLNotFound
	}

	return urls, nil
}
