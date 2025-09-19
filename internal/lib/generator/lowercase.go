package generator

import (
	"math/rand"

	"github.com/nikitaenmi/URLShortener/internal/config"
)

type Lowercase struct {
	cfg config.Generator
}

func NewLowercase(cfg config.Generator) *Lowercase {
	return &Lowercase{
		cfg: cfg,
	}
}

func (g Lowercase) Generate() (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, g.cfg.LengthLetters)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b), nil
}
