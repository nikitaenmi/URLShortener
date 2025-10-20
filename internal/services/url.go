package services

import (
	"context"
	"fmt"

	"github.com/nikitaenmi/URLShortener/internal/domain"
	"github.com/nikitaenmi/URLShortener/internal/lib/generator"
)

type Url struct {
	repo      domain.URLRepo
	generator generator.Generator
}

func NewUrl(repo domain.URLRepo, generator generator.Generator) Url {
	return Url{
		repo:      repo,
		generator: generator,
	}
}

func (s Url) CreateShortURL(ctx context.Context, url domain.Url) (*domain.Url, error) {
	alias, err := s.generator.Generate()
	if err != nil {
		return nil, fmt.Errorf("error generating alias: %w", err)
	}

	url.Alias = alias
	err = s.repo.Create(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed writing url and aliase in database: %w", err)
	}
	return &url, nil
}

func (s Url) Get(ctx context.Context, params domain.URLFilter) (*domain.Url, error) {
	url, err := s.repo.FindById(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("error finding url in database: %w", err)
	}
	return url, nil
}

func (s Url) Update(ctx context.Context, params domain.URLFilter, newURL string) (*domain.Url, error) {
	oldURL, err := s.repo.FindById(ctx, params)
	if err != nil {
		return nil, err
	}

	oldURL.OriginalURL = newURL

	err = s.repo.Update(ctx, oldURL)
	if err != nil {
		return nil, fmt.Errorf("failed to update URL: %w", err)
	}

	return oldURL, nil
}

func (s Url) Delete(ctx context.Context, params domain.URLFilter) error {
	err := s.repo.Delete(ctx, params)
	if err != nil {
		return fmt.Errorf("failed deleting url in database: %w", err)
	}
	return nil
}

func (s Url) List(ctx context.Context, params domain.Paginator) (*domain.UrlList, error) {
	filter := domain.URLFilter{}

	total, err := s.repo.Count(ctx, filter)
	if err != nil {
		return nil, err
	}

	urls, err := s.repo.FindAll(ctx, filter, &params)
	if err != nil {
		return nil, err
	}

	return &domain.UrlList{
		Items: urls,
		Page:  params.Page,
		Limit: params.Limit,
		Total: total,
	}, nil
}
