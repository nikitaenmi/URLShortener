package domain

import (
	"context"
	"net/http"
)

const (
	DefaultLimit = 10
	MaxLimit     = 100
	DefaultPage  = 1
)

type URL struct {
	ID          int
	OriginalURL string
	Alias       string
}

type URLList struct {
	Items []*URL
	Total int
}

type URLRepository interface {
	Create(ctx context.Context, url *URL) error
	Update(ctx context.Context, url *URL) error
	Delete(ctx context.Context, params URLFilter) error
	Count(ctx context.Context, filter URLFilter) (int64, error)
	FindAll(ctx context.Context, filter URLFilter, paginator Paginator) ([]*URL, error)
}

type URLService interface {
	Create(ctx context.Context, url *URL) (*URL, error)
	Get(ctx context.Context, params URLFilter) (*URL, error)
	Update(ctx context.Context, params URLFilter, newURL string) (*URL, error)
	Delete(ctx context.Context, params URLFilter) error
	List(ctx context.Context, params Paginator) (*URLList, error)
}

type URLFilter struct {
	Alias string
	ID    int
}

type Paginator struct {
	Page  int
	Limit int
}

func (p Paginator) GetLimit() int {
	if p.Limit <= 0 || p.Limit > MaxLimit {
		p.Limit = DefaultLimit
	}
	return p.Limit
}

func (p Paginator) GetOffset() int {
	page := p.Page
	if page < 1 {
		page = 1
	}
	limit := p.GetLimit()
	return (page - 1) * limit
}

type Error struct {
	Message string
	Code    int
	Details string
}

var (
	ErrURLNotFound        = &Error{Message: "url not found", Code: http.StatusNotFound}
	ErrInvalidID          = &Error{Message: "invalid ID format", Code: http.StatusBadRequest}
	ErrInvalidRequest     = &Error{Message: "invalid request", Code: http.StatusBadRequest}
	ErrInvalidQueryParams = &Error{Message: "invalid query parameters", Code: http.StatusBadRequest}
	ErrAliasAlreadyExists = &Error{Message: "url already exists", Code: http.StatusConflict}
)

func (e *Error) Error() string {
	return e.Message
}
