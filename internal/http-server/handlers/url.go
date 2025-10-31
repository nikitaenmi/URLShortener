package handlers

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nikitaenmi/URLShortener/internal/config"
	"github.com/nikitaenmi/URLShortener/internal/domain"
	DTO "github.com/nikitaenmi/URLShortener/internal/http-server"
	"github.com/nikitaenmi/URLShortener/internal/lib/logger"
	"github.com/nikitaenmi/URLShortener/internal/services"
)

type URL struct {
	svc services.URL
	log logger.Logger
	cfg config.Server
}

func NewURL(svc services.URL, log logger.Logger, cfg config.Server) URL {
	return URL{
		svc: svc,
		log: log,
		cfg: cfg,
	}
}

func (h *URL) Create(c echo.Context) error {
	ctx := c.Request().Context()
	var req DTO.URLRequest
	if err := c.Bind(&req); err != nil {
		return domain.ErrInvalidRequest
	}
	urlFromUser := req.ToDomain()
	url, err := h.svc.Create(ctx, urlFromUser)
	if err != nil {
		return err
	}
	h.log.Info("link created", slog.String("alias", url.Alias), slog.String("original_url", url.OriginalURL))
	protocol := h.cfg.GetProtocol()
	res := DTO.ToURLItemResponse(protocol, h.cfg.Host, h.cfg.Port, url.Alias)
	return c.JSON(http.StatusCreated, res)
}

func (h *URL) Read(c echo.Context) error {
	ctx := c.Request().Context()
	idStr := c.Param("id")
	id, err := parseID(idStr)
	if err != nil {
		return domain.ErrInvalidID
	}
	url, err := h.svc.Get(ctx, domain.ByID(id))
	if err != nil {
		return err
	}
	h.log.Info("URL retrieved", slog.String("id", idStr))
	return c.JSON(http.StatusOK, url)
}

func (h *URL) Update(c echo.Context) error {
	ctx := c.Request().Context()
	idStr := c.Param("id")
	id, err := parseID(idStr)
	if err != nil {
		return domain.ErrInvalidID
	}
	var req DTO.URLRequest
	if err := c.Bind(&req); err != nil {
		return domain.ErrInvalidRequest
	}
	updatedURL, err := h.svc.Update(ctx, domain.ByID(id), req.OriginalURL)
	if err != nil {
		return err
	}
	h.log.Info("url updated", slog.Int("id", id), slog.String("original_url", req.OriginalURL))
	return c.JSON(http.StatusOK, updatedURL)
}

func (h *URL) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	idStr := c.Param("id")
	id, err := parseID(idStr)
	if err != nil {
		return domain.ErrInvalidID
	}
	if err := h.svc.Delete(ctx, domain.ByID(id)); err != nil {
		return err
	}
	h.log.Info("url deleted", slog.Int("id", id))
	return c.NoContent(http.StatusNoContent)
}

func (h *URL) List(c echo.Context) error {
	ctx := c.Request().Context()
	var req DTO.Paginator
	if err := c.Bind(&req); err != nil {
		return domain.ErrInvalidQueryParams
	}

	req.ValidateAndSetDefaults()

	paginator := req.ToDomain()
	urlList, err := h.svc.List(ctx, paginator)
	if err != nil {
		return err
	}

	res := DTO.ToURLListResponse(urlList)
	h.log.Info("URLs retrieved")
	return c.JSON(http.StatusOK, res)
}

func (h *URL) Redirect(c echo.Context) error {
	ctx := c.Request().Context()
	alias := c.Param("alias")
	url, err := h.svc.Get(ctx, domain.ByAlias(alias))
	if err != nil {
		return err
	}
	h.log.Info("Redirection", slog.String("alias", alias), slog.String("original_url", url.OriginalURL))
	return c.Redirect(http.StatusMovedPermanently, url.OriginalURL)
}
