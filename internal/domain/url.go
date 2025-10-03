package domain

import "context"

type Url struct {
	ID          int
	OriginalURL string
	Alias       string
}

type URLFilter struct {
	Alias string
	ID    int
}

type Paginator struct {
	Limit int
	Page  int
	Total int64
}

type URLRepo interface {
	Create(ctx context.Context, url Url) error
	FindById(ctx context.Context, params URLFilter) (*Url, error)
	Delete(ctx context.Context, params URLFilter) error
	List(ctx context.Context, params Paginator) ([]*Url, Paginator, error)
	Update(ctx context.Context, url *Url) error
	Count(ctx context.Context, filter URLFilter) (int64, error)
	FindAll(ctx context.Context, filter URLFilter, paginator *Paginator) ([]*Url, error)
}
