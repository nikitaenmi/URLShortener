package httpserver

import (
	"github.com/nikitaenmi/URLShortener/internal/domain"
)

// Единый request для создания и обновления URL
type UrlRequest struct {
	OriginalURL string `json:"original_url"`
	Alias       string `json:"alias"`
}

type UrlItemResponse struct {
	ShortURL string `json:"short_url"`
}

type UrlListRequest struct {
	Page  int `query:"page"`
	Limit int `query:"limit"`
}

type UrlListResponse struct {
	URLs  []*domain.Url `json:"urls"`
	Total int           `json:"total"`
}
