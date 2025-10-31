package services

import (
	"context"
	"fmt"

	"github.com/nikitaenmi/URLShortener/internal/domain"
	"github.com/nikitaenmi/URLShortener/internal/lib/generator"
)

type URL struct {
	repo      domain.URLRepo
	generator generator.Generator
}

func NewURL(repo domain.URLRepo, generator generator.Generator) URL {
	return URL{
		repo:      repo,
		generator: generator,
	}
}

func (s URL) Create(ctx context.Context, url domain.URL) (*domain.URL, error) {
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

func (s URL) Get(ctx context.Context, params domain.URLFilter) (*domain.URL, error) {
	urls, err := s.repo.FindAll(ctx, params, nil)
	if err != nil {
		return nil, fmt.Errorf("error finding url in database: %w", err)
	}
	return urls[0], nil
}

func (s URL) Update(ctx context.Context, params domain.URLFilter, newURL string) (*domain.URL, error) {
	urls, err := s.repo.FindAll(ctx, params, nil)
	if err != nil {
		return nil, err
	}

	oldURL := urls[0]
	oldURL.OriginalURL = newURL

	err = s.repo.Update(ctx, oldURL)
	if err != nil {
		return nil, fmt.Errorf("failed to update URL: %w", err)
	}

	return oldURL, nil
}

func (s URL) Delete(ctx context.Context, params domain.URLFilter) error {
	err := s.repo.Delete(ctx, params)
	if err != nil {
		return fmt.Errorf("failed deleting url in database: %w", err)
	}
	return nil
}

func (s URL) List(ctx context.Context, params domain.Paginator) (*domain.URLList, error) {
	filter := domain.URLFilter{}

	total, err := s.repo.Count(ctx, filter)
	if err != nil {
		return nil, err
	}

	urls, err := s.repo.FindAll(ctx, filter, &params)
	if err != nil {
		return nil, err
	}

	return &domain.URLList{
		Items: urls,
		Total: total,
	}, nil
}
