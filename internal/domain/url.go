package domain

import (
	"context"
	"errors"
)

type Url struct {
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
	Limit int
	Page  int
	Total int64
}

type URLRepo interface {
	Create(ctx context.Context, url Url) error
	Update(ctx context.Context, url *Url) error
	Delete(ctx context.Context, params URLFilter) error
	Count(ctx context.Context, filter URLFilter) (int64, error)
	FindById(ctx context.Context, params URLFilter) (*Url, error)
	FindAll(ctx context.Context, filter URLFilter, paginator *Paginator) ([]*Url, error)
}

var (
	ErrURLNotFound = errors.New("url not found")
)

type UrlList struct {
	Items []*Url
	Page  int
	Limit int
	Total int64
}
