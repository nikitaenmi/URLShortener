package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
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

func (h *Url) RedirectURL(c echo.Context) error {
	ctx := c.Request().Context()
	alias := c.Param("alias")
	params := domain.URLFilter{
		Alias: alias,
	}
	

	url, err := h.svc.Redirect(ctx,params)
	if err != nil {
		h.log.Error("Link not found", slog.String("alias", alias), slog.String("error", err.Error()))
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Link not found"})
	}

	h.log.Info("Redirection", slog.String("alias", alias), slog.String("original_url", url.OriginalURL))
	return c.Redirect(http.StatusMovedPermanently, url.OriginalURL)
}

func (h *Url) ShortenerURL(c echo.Context) error {
	ctx := c.Request().Context()
	var request struct {
		OriginalURL string `json:"url"`
	}
	

	if err := c.Bind(&request); err != nil {
		h.log.Error("Invalid request", slog.String("error", err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	url := domain.Url{
		OriginalURL: request.OriginalURL,
	}

	alias, err := h.svc.Shortener(ctx,url)
	if err != nil {
		h.log.Error("Shortener failed", slog.String("error", err.Error()))
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	h.log.Info("Link created successfully!", slog.String("alias", alias))

	response := map[string]string{
		"short_url": fmt.Sprintf("http://%s:%s/%s", h.cfg.Host, h.cfg.Port, alias),
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Url) DeleteURL(c echo.Context) error {
	ctx := c.Request().Context()
	alias := c.Param("alias")
	params := domain.URLFilter{
		Alias: alias,
	}
	err := h.svc.Delete(ctx, params)
	
	if err != nil {
		h.log.Error("Failed to delete URL", slog.String("alias", alias), slog.String("error", err.Error()))
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete URL"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "URL deleted successfully"})
}
