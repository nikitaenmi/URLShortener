package shortener

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nikitaenmi/URLShortener/internal/config"
	"github.com/nikitaenmi/URLShortener/internal/services"
)

type Creater interface {
	Create(URL, alias string) error
}

type Logger interface {
	Info(msg string, args ...any)
	Error(msg string, args ...any)
}

func ShortenerURL(w http.ResponseWriter, r *http.Request, repo Creater, cfg config.Server, log Logger) {
	var request struct {
		URL string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		log.Error("Invalid request", err)
		return
	}

	alias, err := services.Shortener(request.URL, repo)
	if err != nil {
		log.Error("Shortener failed", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Info("Link created successfully!", "alias", alias)

	response := map[string]string{
		"short_url": fmt.Sprintf("http://%s:%s/%s", cfg.Host, cfg.Port, alias),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
