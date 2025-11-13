package handlers

import (
	"strconv"

	"github.com/nikitaenmi/URLShortener/internal/domain"
)

func parseID(idStr string) (int, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, domain.ErrInvalidID
	}
	return id, nil
}
