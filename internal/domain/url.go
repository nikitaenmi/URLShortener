package domain

import "context"

type Url struct {
	ID          int
	OriginalURL string
	Alias       string
}

type URLFilter struct {
	Alias string
}

type UrlRepo interface {
	Create(ctx context.Context, url Url) error
	URLFind(ctx context.Context, params URLFilter) (*Url, error)
	Delete(ctx context.Context, params URLFilter) error
}
