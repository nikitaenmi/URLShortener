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

func (s Url) Shortener(ctx context.Context, url domain.Url) (string, error) {
	alias, err := s.generator.Generate()
	if err != nil {
		return "", fmt.Errorf("error generating alias: %w", err)
	}

	url.Alias = alias
	err = s.repo.Create(ctx, url)
	if err != nil {
		return "", fmt.Errorf("failed writing url and aliase in database: %w", err)
	}
	return alias, nil
}

func (s Url) Get(ctx context.Context, params domain.URLFilter) (*domain.Url, error) {
	url, err := s.repo.Find(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("error finding url in database: %w", err)
	}
	return url, nil
}

func (s Url) Update(ctx context.Context, params domain.URLFilter, newURL string) (*domain.Url, error) {
	existingURL, err := s.repo.Find(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("URL not found: %w", err)
	}

	existingURL.OriginalURL = newURL

	err = s.repo.Update(ctx, existingURL)
	if err != nil {
		return nil, fmt.Errorf("failed to update URL: %w", err)
	}

	return existingURL, nil
}

func (s Url) Delete(ctx context.Context, params domain.URLFilter) error {
	err := s.repo.Delete(ctx, params)
	if err != nil {
		return fmt.Errorf("failed deleting url in database: %w", err)
	}
	return nil
}

func (s Url) List(ctx context.Context, params domain.Paginator) ([]*domain.Url, domain.Paginator, error) {
	urls, total, err := s.repo.List(ctx, params)
	if err != nil {
		return nil, params, fmt.Errorf("failed to get URLs list: %w", err)
	}

	return urls, total, nil
}
