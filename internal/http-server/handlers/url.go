package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

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

// CRUDL - OPERATION

func (h *Url) Create(c echo.Context) error {
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

	alias, err := h.svc.Shortener(ctx, url)
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

func (h *Url) Get(c echo.Context) error {
	ctx := c.Request().Context()
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		h.log.Error("Invalid ID format", slog.String("id", idString), slog.String("error", err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
	}

	params := domain.URLFilter{
		ID: id,
	}

	url, err := h.svc.Get(ctx, params)
	if err != nil {
		h.log.Error("URL not found", slog.String("id", idString), slog.String("error", err.Error()))
		return c.JSON(http.StatusNotFound, map[string]string{"error": "URL not found"})
	}

	h.log.Info("URL retrieved successfully", slog.String("id", idString))
	return c.JSON(http.StatusOK, url)
}

type URLListResponse struct {
	URLs  []*domain.Url `json:"urls"`
	Total int           `json:"total"`
	Page  int           `json:"page"`
	Limit int           `json:"limit"`
}

func (h *Url) Put(c echo.Context) error {
	ctx := c.Request().Context()
	idString := c.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		h.log.Error("Invalid ID format", slog.String("id", idString), slog.String("error", err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
	}

	var request struct {
		OriginalURL string `json:"url"`
	}

	if err := c.Bind(&request); err != nil {
		h.log.Error("Invalid request", slog.String("error", err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	params := domain.URLFilter{
		ID: id,
	}

	updatedURL, err := h.svc.Update(ctx, params, request.OriginalURL)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			h.log.Error("URL not found", slog.String("id", idString), slog.String("error", err.Error()))
			return c.JSON(http.StatusNotFound, map[string]string{"error": "URL not found"})
		}
		h.log.Error("Failed to update URL", slog.String("id", idString), slog.String("error", err.Error()))
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update URL"})
	}

	h.log.Info("URL updated successfully", slog.String("id", idString))
	return c.JSON(http.StatusOK, updatedURL)
}

func (h *Url) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		h.log.Error("Invalid ID format", slog.String("id", idString), slog.String("error", err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
	}
	params := domain.URLFilter{
		ID: id,
	}
	err = h.svc.Delete(ctx, params)
	if err != nil {
		h.log.Error("Failed to delete URL", slog.String("id", idString), slog.String("error", err.Error()))
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete URL"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "URL deleted successfully"})
}

func (h *Url) List(c echo.Context) error {
	ctx := c.Request().Context()

	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}

	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit < 1 || limit > 50 {
		limit = 10
	}

	urls, total, err := h.svc.List(ctx, page, limit)
	if err != nil {
		h.log.Error("Failed to list URLs", slog.String("error", err.Error()))
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to get URLs",
		})
	}

	response := URLListResponse{
		URLs:  urls,
		Total: total,
		Page:  page,
		Limit: limit,
	}

	return c.JSON(http.StatusOK, response)
}

// REDIRECT - OPETATION

func (h *Url) Redirect(c echo.Context) error {
	ctx := c.Request().Context()

	alias := c.Param("alias")
	params := domain.URLFilter{
		Alias: alias,
	}

	url, err := h.svc.Get(ctx, params)
	if err != nil {
		h.log.Error("Link not found", slog.String("alias", alias), slog.String("error", err.Error()))
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Link not found"})
	}

	h.log.Info("Redirection", slog.String("alias", alias), slog.String("original_url", url.OriginalURL))
	return c.Redirect(http.StatusMovedPermanently, url.OriginalURL)
}
