package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nikitaenmi/URLShortener/internal/config"
	"github.com/nikitaenmi/URLShortener/internal/domain"
	"github.com/nikitaenmi/URLShortener/internal/lib/logger"
	"github.com/nikitaenmi/URLShortener/internal/services"
)

type Url struct {
	svc services.Url
	log logger.Logger
	cfg config.Server
}

func NewUrl(svc services.Url, log logger.Logger, cfg config.Server) Url {
	return Url{
		svc: svc,
		log: log,
		cfg: cfg,
	}
}

func (h *Url) RedirectURL(w http.ResponseWriter, r *http.Request) {
	alias := r.URL.Path[1:]
	params := domain.URLFilter{
		Alias: alias,
	}

	url, err := h.svc.Redirect(params)
	if err != nil {
		h.log.Error("Link not found")
		http.Error(w, "Link not found", http.StatusNotFound)
		return
	}

	h.log.Info("Redirection")
	http.Redirect(w, r, url.OriginalURL, http.StatusMovedPermanently)
}

func (h *Url) ShortenerURL(w http.ResponseWriter, r *http.Request) {
	var request struct {
		OriginalURL string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		h.log.Error("Invalid request", err)
		return
	}

	url := domain.Url{
		OriginalURL: request.OriginalURL,
	}

	alias, err := h.svc.Shortener(url)
	if err != nil {
		h.log.Error("Shortener failed", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	h.log.Info("Link created successfully!", "alias", alias)

	response := map[string]string{
		"short_url": fmt.Sprintf("http://%s:%s/%s", h.cfg.Host, h.cfg.Port, alias),
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.log.Error("Shortener failed", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
