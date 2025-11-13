package httpserver

import (
	"github.com/nikitaenmi/URLShortener/internal/domain"
)

// Единый request для создания и обновления URL
type URLRequest struct {
	OriginalURL string `json:"original_url"`
	Alias       string `json:"alias"`
}

type URLItemResponse struct {
	ShortURL string `json:"short_url"`
}

type Paginator struct {
	Page  int `query:"page"`
	Limit int `query:"limit"`
}

type URLListResponse struct {
	URLs  []*domain.URL `json:"urls"`
	Total int           `json:"total"`
}
