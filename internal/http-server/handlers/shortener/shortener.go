package shortener

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nikitaenmi/URLShortener/internal/config"
	"github.com/teris-io/shortid"
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
	
	alias, err := shortid.Generate()
	if err != nil {
		http.Error(w, "Error generating alias", http.StatusInternalServerError)
		log.Error("Error generating alias: %v\n", err)
		return
	}

	err = repo.Create(request.URL, alias)
	if err != nil {
		http.Error(w, "Creating error", http.StatusInternalServerError)
		log.Error("Error generating alias: %v\n", err)
		return
	}

	log.Info("Link created successfully!", "alias", alias)

	response := map[string]string{
		"short_url": fmt.Sprintf("http://%s:%s/%s", cfg.Host, cfg.Port, alias),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
