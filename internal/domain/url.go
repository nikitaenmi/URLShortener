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

type URLRepo interface {
	Create(ctx context.Context, url Url) error
	Find(ctx context.Context, params URLFilter) (*Url, error)
	Delete(ctx context.Context, params URLFilter) error
	List(ctx context.Context, page, limit int) ([]*Url, int, error)
	Update(ctx context.Context, url *Url) error
}
