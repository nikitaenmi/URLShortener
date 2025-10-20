package handlers

import (
	"strconv"

	"github.com/nikitaenmi/URLShortener/internal/lib/logger"
)

func ParseID(idStr string, log logger.Logger) (int, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Warn("Invalid ID format", "id", idStr)
		return 0, err
	}
	return id, nil
}
