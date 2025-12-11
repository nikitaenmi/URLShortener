package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/nikitaenmi/URLShortener/internal/domain"
	"github.com/nikitaenmi/URLShortener/internal/lib/generator"
)

type URL struct {
	repo domain.URLRepository
	gen  generator.Generator
}

func NewURL(repo domain.URLRepository, generator generator.Generator) URL {
	return URL{
		repo: repo,
		gen:  generator,
	}
}

func (s URL) Create(ctx context.Context, url *domain.URL) (*domain.URL, error) {
	alias, err := s.gen.Generate()
	if err != nil {
		return nil, fmt.Errorf("error generating alias: %w", err)
	}

	url.Alias = alias
	err = s.repo.Create(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed writing url and aliase in database: %w", err)
	}
	return url, nil
}

func (s URL) Get(ctx context.Context, params domain.URLFilter) (*domain.URL, error) {
	urls, err := s.repo.FindAll(ctx, params, domain.Paginator{})
	if err != nil {
		if _, ok := err.(*domain.Error); ok {
			return nil, err
		}
		return nil, fmt.Errorf("failed finding url in database: %w", err)
	}

	if len(urls) == 0 {
		return nil, domain.ErrURLNotFound
	}

	return urls[0], nil
}

func (s URL) Update(ctx context.Context, params domain.URLFilter, newURL string) (*domain.URL, error) {
	urls, err := s.repo.FindAll(ctx, params, domain.Paginator{})
	if err != nil {
		if _, ok := err.(*domain.Error); ok {
			return nil, err
		}
		return nil, fmt.Errorf("failed finding url in database: %w", err)
	}

	if len(urls) == 0 {
		return nil, fmt.Errorf("error record url in list: %w", errors.New("list is empty"))
	}

	oldURL := urls[0]
	oldURL.OriginalURL = newURL

	err = s.repo.Update(ctx, oldURL)
	if err != nil {
		if _, ok := err.(*domain.Error); ok {
			return nil, err
		}
		return nil, fmt.Errorf("failed updating URL: %w", err)
	}

	return oldURL, nil
}

func (s URL) Delete(ctx context.Context, params domain.URLFilter) error {
	err := s.repo.Delete(ctx, params)
	if err != nil {
		if _, ok := err.(*domain.Error); ok {
			return err
		}
		return fmt.Errorf("failed deleting url in database: %w", err)
	}
	return nil
}

func (s URL) List(ctx context.Context, params domain.Paginator) (*domain.URLList, error) {
	filter := domain.URLFilter{}

	total, err := s.repo.Count(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to counting: %w", err)
	}

	urls, err := s.repo.FindAll(ctx, filter, params)
	if err != nil {
		if _, ok := err.(*domain.Error); ok {
			return nil, err
		}
		return nil, fmt.Errorf("failed finding url in database: %w", err)
	}
	return &domain.URLList{
		Items: urls,
		Total: int(total),
	}, nil
}
