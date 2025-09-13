package config

type Generator struct {
	LengthLetters int    `env:"LENGTH_LETTERS"`
	Type          string `env:"TYPE"`
}
