package services

import (
	"fmt"

	"github.com/nikitaenmi/URLShortener/internal/domain"
	"github.com/teris-io/shortid"
)

type Creater interface {
	Create(URL, alias string) error
}

type URLFinder interface {
	URLFind(alias string) (*domain.Link, error)
}

func Shortener(url string, repo Creater) (string, error) {
	alias, err := shortid.Generate()
	if err != nil {
		return "", fmt.Errorf("error generating alias: %w", err)
	}

	err = repo.Create(url, alias)
	if err != nil {
		return "", fmt.Errorf("failed writing url and aliase in database: %w", err)
	}

	return alias, nil
}

func Redirect(alias string, repo URLFinder) (*domain.Link, error) {
	link, err := repo.URLFind(alias)
	if err != nil {
		return nil, fmt.Errorf("error finding url in database: %w", err)
	}
	return link, nil
}
