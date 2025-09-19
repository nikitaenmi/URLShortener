package generator

import (
	"math/rand"

	"github.com/nikitaenmi/URLShortener/internal/config"
)

type Digit struct {
	cfg config.Generator
}

func NewDigit(cfg config.Generator) *Digit {
	return &Digit{
		cfg: cfg,
	}
}

func (g Digit) Generate() (string, error) {
	const digits = "0123456789"
	b := make([]byte, g.cfg.LengthLetters)
	for i := range b {
		b[i] = digits[rand.Intn(len(digits))]
	}

	return string(b), nil
}
