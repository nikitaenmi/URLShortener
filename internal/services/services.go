package services

import (
	"fmt"

	"github.com/teris-io/shortid"
)

type Creater interface {
	Create(URL, alias string) error
}

func Shortener(url string, repo Creater) (string, error) {
	alias, err := shortid.Generate()
	if err != nil {
		return "", fmt.Errorf("failed to generate alias: %w", err)
	}

	if err := repo.Create(url, alias); err != nil {
		return "", fmt.Errorf("failed to save URL: %w", err)
	}

	return alias, nil
}
