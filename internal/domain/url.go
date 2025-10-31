package domain

import (
	"context"
	"errors"

	"github.com/nikitaenmi/URLShortener/internal/constants"
)

type URL struct {
	ID          int
	OriginalURL string
	Alias       string
}

type URLFilter struct {
	Alias string
	ID    int
}

func ByID(id int) URLFilter {
	return URLFilter{ID: id}
}

func ByAlias(alias string) URLFilter {
	return URLFilter{Alias: alias}
}

type Paginator struct {
	Page  int
	Limit int
}

func (p Paginator) GetLimit() int {
	if p.Limit <= 0 || p.Limit > constants.MaxLimit {
		p.Limit = constants.DefaultLimit
	}
	return p.Limit
}
func (p Paginator) GetPage() int {
	if p.Page <= 0 {
		p.Page = constants.DefaultPage
	}
	return p.Page
}

func (p Paginator) GetOffset() int {
	return (p.Page - 1) * p.Limit
}

type URLRepo interface {
	Create(ctx context.Context, url URL) error
	Update(ctx context.Context, url *URL) error
	Delete(ctx context.Context, params URLFilter) error
	Count(ctx context.Context, filter URLFilter) (int64, error)
	FindAll(ctx context.Context, filter URLFilter, paginator *Paginator) ([]*URL, error)
}

var (
	ErrURLNotFound        = errors.New("url not found")
	ErrInvalidRequest     = errors.New("invalid request")
	ErrInvalidID          = errors.New("invalid ID format")
	ErrInvalidQueryParams = errors.New("invalid query parameters")
)

type URLList struct {
	Items []*URL
	Total int64
}
