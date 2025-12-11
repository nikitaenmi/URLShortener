package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nikitaenmi/URLShortener/internal/config"
	"github.com/nikitaenmi/URLShortener/internal/domain"
	server "github.com/nikitaenmi/URLShortener/internal/http-server"
	"github.com/nikitaenmi/URLShortener/internal/lib/logger"
	"github.com/nikitaenmi/URLShortener/internal/services"
)

const (
	urlColumnID       = "id"
	urlOriginalColumn = "original_url"
	urlColumnAlias    = "alias"
)

type URL struct {
	svc domain.URLService
	log logger.Logger
	cfg *config.App
}

func NewURL(svc services.URL, log logger.Logger, cfg *config.App) *URL {
	return &URL{
		svc: svc,
		log: log,
		cfg: cfg,
	}
}

type getURLParams struct {
	ID int `param:"id"`
}

func (h *URL) Create(c echo.Context) error {
	ctx := c.Request().Context()

	var req server.URLRequest
	if err := c.Bind(&req); err != nil {
		return &domain.Error{
			Message: domain.ErrInvalidRequest.Message,
			Code:    domain.ErrInvalidRequest.Code,
			Details: err.Error(),
		}
	}

	if req.OriginalURL == "" {
		return domain.ErrInvalidRequest
	}

	inputURL := req.ToDomain()
	url, err := h.svc.Create(ctx, inputURL)
	if err != nil {
		return err
	}
	h.log.Info(ctx, "link created",
		urlColumnAlias, url.Alias,
		urlOriginalColumn, url.OriginalURL,
	)
	res := server.ToURLItemResponse(url)
	return c.JSON(http.StatusCreated, res)
}

func (h *URL) Read(c echo.Context) error {
	ctx := c.Request().Context()

	params := new(getURLParams)
	if err := c.Bind(params); err != nil {
		return &domain.Error{
			Message: domain.ErrInvalidID.Message,
			Code:    domain.ErrInvalidID.Code,
			Details: err.Error(),
		}
	}

	url, err := h.svc.Get(ctx, domain.URLFilter{ID: params.ID})
	if err != nil {
		return err
	}
	h.log.Info(ctx, "url retrieved", urlColumnID, params.ID)
	return c.JSON(http.StatusOK, url)
}

func (h *URL) Update(c echo.Context) error {
	ctx := c.Request().Context()

	params := new(getURLParams)
	if err := c.Bind(params); err != nil {
		return &domain.Error{
			Message: domain.ErrInvalidID.Message,
			Code:    domain.ErrInvalidID.Code,
			Details: err.Error(),
		}
	}

	var req server.URLRequest
	if err := c.Bind(&req); err != nil {
		return &domain.Error{
			Message: domain.ErrInvalidRequest.Message,
			Code:    domain.ErrInvalidRequest.Code,
			Details: err.Error(),
		}
	}

	if req.OriginalURL == "" {
		return domain.ErrInvalidRequest
	}

	updatedURL, err := h.svc.Update(ctx, domain.URLFilter{ID: params.ID}, req.OriginalURL)
	if err != nil {
		return err
	}
	h.log.Info(ctx, "url updated",
		urlColumnID, params.ID,
		urlOriginalColumn, req.OriginalURL,
	)
	return c.JSON(http.StatusOK, updatedURL)
}

func (h *URL) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	params := new(getURLParams)
	if err := c.Bind(params); err != nil {
		return &domain.Error{
			Message: domain.ErrInvalidID.Message,
			Code:    domain.ErrInvalidID.Code,
			Details: err.Error(),
		}
	}
	if err := h.svc.Delete(ctx, domain.URLFilter{ID: params.ID}); err != nil {
		return err
	}
	h.log.Info(ctx, "url deleted", urlColumnID, params.ID)
	return c.NoContent(http.StatusNoContent)
}

func (h *URL) List(c echo.Context) error {
	ctx := c.Request().Context()

	var req server.Paginator
	if err := c.Bind(&req); err != nil {
		return &domain.Error{
			Message: domain.ErrInvalidQueryParams.Message,
			Code:    domain.ErrInvalidQueryParams.Code,
			Details: err.Error(),
		}
	}

	paginator := req.ToDomain()
	urlList, err := h.svc.List(ctx, paginator)
	if err != nil {
		return err
	}

	res := server.ToURLListResponse(urlList)
	h.log.Info(ctx, "urls retrieved")
	return c.JSON(http.StatusOK, res)
}

func (h *URL) Redirect(c echo.Context) error {
	ctx := c.Request().Context()

	alias := c.Param("alias")
	if alias == "" {
		return domain.ErrInvalidRequest
	}

	url, err := h.svc.Get(ctx, domain.URLFilter{Alias: alias})
	if err != nil {
		return err
	}

	h.log.Info(ctx, "redirection",
		urlColumnAlias, alias,
		urlOriginalColumn, url.OriginalURL,
	)
	return c.Redirect(http.StatusMovedPermanently, url.OriginalURL)
}
