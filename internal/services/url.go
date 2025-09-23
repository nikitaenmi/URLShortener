package services

import (
	"context"
	"fmt"

	"github.com/nikitaenmi/URLShortener/internal/domain"
	"github.com/nikitaenmi/URLShortener/internal/lib/generator"
)

type Url struct {
	repo      domain.UrlRepo
	generator generator.Generator
}

func NewUrl(repo domain.UrlRepo, generator generator.Generator) Url {
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
	fmt.Println(url)
	err = s.repo.Create(ctx, url)
	if err != nil {
		return "", fmt.Errorf("failed writing url and aliase in database: %w", err)
	}
	return alias, nil
}

func (s Url) Redirect(ctx context.Context, params domain.URLFilter) (*domain.Url, error) {
	url, err := s.repo.URLFind(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("error finding url in database: %w", err)
	}
	return url, nil
}

func (s Url) Delete(ctx context.Context, params domain.URLFilter) error {
	err := s.repo.Delete(ctx, params)
	if err != nil {
		return fmt.Errorf("failed deleting url in database: %w", err)
	}
	return nil
}


func (s Url) GetByID(ctx context.Context, params domain.URLFilter) (*domain.Url, error) {
	url, err := s.repo.URLFind(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("error finding url in database: %w", err)
	}
	return url, nil
}
