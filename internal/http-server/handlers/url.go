package handlers

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nikitaenmi/URLShortener/internal/config"
	"github.com/nikitaenmi/URLShortener/internal/domain"
	dto "github.com/nikitaenmi/URLShortener/internal/http-server"
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

func (h *Url) Create(c echo.Context) error {
	ctx := c.Request().Context()

	var req dto.UrlRequest

	if err := c.Bind(&req); err != nil {
		h.log.Error("Invalid request", slog.String("error", err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	url := req.ToDomain()

	shortURL, err := h.svc.CreateShortURL(ctx, url)
	if err != nil {
		return err
	}

	h.log.Info("link created", slog.String("alias", shortURL.Alias), slog.String("original_url", url.OriginalURL))
	res := dto.ToUrlItemResponse(shortURL.Alias, h.cfg.Host, h.cfg.Port)
	return c.JSON(http.StatusCreated, res)
}

func (h *Url) Read(c echo.Context) error {
	ctx := c.Request().Context()

	idStr := c.Param("id")

	id, err := ParseID(idStr, h.log)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
	}

	url, err := h.svc.Get(ctx, domain.ByID(id))
	if err != nil {
		return err
	}

	h.log.Info("URL retrieved", slog.String("id", idStr))
	return c.JSON(http.StatusOK, url)
}

func (h *Url) Update(c echo.Context) error {
	ctx := c.Request().Context()

	idStr := c.Param("id")

	id, err := ParseID(idStr, h.log)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
	}

	var req dto.UrlRequest

	if err := c.Bind(&req); err != nil {
		h.log.Error("Invalid request", slog.String("error", err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	updatedURL, err := h.svc.Update(ctx, domain.ByID(id), req.OriginalURL)
	if err != nil {
		return err
	}

	h.log.Info("url updated", slog.Int("id", id), slog.String("original_url", req.OriginalURL))
	return c.JSON(http.StatusOK, updatedURL)
}

func (h *Url) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	idStr := c.Param("id")

	id, err := ParseID(idStr, h.log)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
	}

	if err := h.svc.Delete(ctx, domain.ByID(id)); err != nil {
		return err
	}
	h.log.Info("url deleted", slog.Int("id", id))

	return c.NoContent(http.StatusNoContent)
}

func (h *Url) List(c echo.Context) error {
	ctx := c.Request().Context()

	var req dto.UrlListRequest
	if err := c.Bind(&req); err != nil {
		h.log.Warn("Invalid query parameters", "error", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid query parameters",
		})
	}

	req.ValidateAndSetDefaults()
	par := req.ToDomain()

	urlList, err := h.svc.List(ctx, par)
	if err != nil {
		return err
	}

	res := dto.ToUrlListResponse(urlList)
	h.log.Info("URLs retrieved")
	return c.JSON(http.StatusOK, res)
}

func (h *Url) Redirect(c echo.Context) error {
	ctx := c.Request().Context()

	alias := c.Param("alias")

	url, err := h.svc.Get(ctx, domain.ByAlias(alias))
	if err != nil {
		return err
	}

	h.log.Info("Redirection", slog.String("alias", alias), slog.String("original_url", url.OriginalURL))
	return c.Redirect(http.StatusMovedPermanently, url.OriginalURL)
}
