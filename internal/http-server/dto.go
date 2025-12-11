package httpserver

import (
	"github.com/nikitaenmi/URLShortener/internal/domain"
)

type URLRequest struct {
	OriginalURL string `json:"original_url"`
	Alias       string `json:"alias"`
}

type URLItemResponse struct {
	ID          int    `json:"id"`
	OriginalURL string `json:"original_url"`
	Alias       string `json:"alias"`
}

type URLListResponse struct {
	URLs  []*domain.URL `json:"urls"`
	Total int           `json:"total"`
}

type Paginator struct {
	Page  int `query:"page"`
	Limit int `query:"limit"`
}
