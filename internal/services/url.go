package services

import (
	"fmt"

	"github.com/nikitaenmi/URLShortener/internal/domain"
	"github.com/teris-io/shortid"
)

type Url struct {
	repo domain.UrlRepo
}

func NewUrl(repo domain.UrlRepo) Url {
	return Url{
		repo: repo,
	}
}

func (s *Url) Shortener(url domain.Url) (string, error) {
	alias, err := shortid.Generate()
	if err != nil {
		return "", fmt.Errorf("error generating alias: %w", err)
	}

	url.Alias = alias
	fmt.Println(url)
	err = s.repo.Create(url)
	if err != nil {
		return "", fmt.Errorf("failed writing url and aliase in database: %w", err)
	}

	return alias, nil
}

func (s *Url) Redirect(params domain.URLFilter) (*domain.Url, error) {
	url, err := s.repo.URLFind(params)
	if err != nil {
		return nil, fmt.Errorf("error finding url in database: %w", err)
	}
	return url, nil
}
