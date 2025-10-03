package httpserver

import (
	"github.com/nikitaenmi/URLShortener/internal/domain"
)

type Response struct {
	URLs  []*domain.Url `json:"urls"`
	Total int           `json:"total"`
	Page  int           `json:"page"`
	Limit int           `json:"limit"`
}

type Request struct {
	Page  int `query:"page" validate:"min=1"`
	Limit int `query:"limit" validate:"min=1,max=50"`
}
