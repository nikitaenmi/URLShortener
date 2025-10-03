package httpserver

import (
	"github.com/nikitaenmi/URLShortener/internal/domain"
)

type CreateRequest struct {
	OriginalURL string `json:"original_url"`
}

type CreateResponse struct {
	ShortURL string `json:"short_url"`
}

type PutRequest struct {
	OriginalURL string `json:"url"`
}
type PutResponse struct {
	ShortURL string `json:"short_url"`
}

type ListResponse struct {
	URLs  []*domain.Url `json:"urls"`
	Total int           `json:"total"`
	Page  int           `json:"page"`
	Limit int           `json:"limit"`
}

type ListRequest struct {
	Page  int `query:"page" validate:"min=1"`
	Limit int `query:"limit" validate:"min=1,max=50"`
}
