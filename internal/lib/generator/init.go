package generator

import (
	"errors"

	"github.com/nikitaenmi/URLShortener/internal/config"
)

const (
	DigitType     = "digit"
	LowercaseType = "lowercase"
)

var ErrUnknownGeneratorType = errors.New("unknown generator type")

type Generator interface {
	Generate() (string, error)
}

func New(cfg config.Generator) (Generator, error) {
	var g Generator

	switch cfg.Type {
	case DigitType:
		g = NewDigit(cfg)
	case LowercaseType:
		g = NewLowercase(cfg)
	default:
		return g, ErrUnknownGeneratorType
	}

	return g, nil
}
