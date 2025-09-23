package domain

import "context"

type Url struct {
	ID          string
	OriginalURL string
	Alias       string
}

type URLFilter struct {
	Alias string
	ID    string
}

type UrlRepo interface {
	Create(ctx context.Context, url Url) error
	URLFind(ctx context.Context, id URLFilter) (*Url, error)
	Delete(ctx context.Context, id URLFilter) error
}
