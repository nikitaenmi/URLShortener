package httpserver

import (
	"github.com/nikitaenmi/URLShortener/internal/domain"
)

const (
	DefaultLimit = 10
	MaxLimit     = 100
	DefaultPage  = 1
)

func (r *URLRequest) ToDomain() *domain.URL {
	return &domain.URL{
		OriginalURL: r.OriginalURL,
	}
}

func ToURLItemResponse(url *domain.URL) URLItemResponse {
	return URLItemResponse{
		ID:          url.ID,
		OriginalURL: url.OriginalURL,
		Alias:       url.Alias,
	}
}

func (r *Paginator) ToDomain() domain.Paginator {
	if r.Page < 1 {
		r.Page = DefaultPage
	}
	if r.Limit < 1 || r.Limit > MaxLimit {
		r.Limit = DefaultLimit
	}
	return domain.Paginator{
		Page:  r.Page,
		Limit: r.Limit,
	}
}

func ToURLListResponse(urlList *domain.URLList) URLListResponse {
	return URLListResponse{
		URLs:  urlList.Items,
		Total: urlList.Total,
	}
}
