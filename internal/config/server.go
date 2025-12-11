package config

import (
	"time"

	"github.com/caarlos0/env/v11"
)

const (
	HTTPProtocol  = "http"
	HTTPSProtocol = "https"
)

type Server struct {
	Host         string        `env:"SERVER_HOST"`
	Port         string        `env:"SERVER_PORT"`
	ReadTimeout  time.Duration `env:"SERVER_READ_TIMEOUT_SECOND"`
	WriteTimeout time.Duration `env:"SERVER_WRITE_TIMEOUT_SECOND"`
	IdleTimeout  time.Duration `env:"SERVER_IDLETIMEOUT_S"`
	Secure       bool          `env:"SERVER_SECURE_ENABLE"`
}

func (s *Server) GetProtocol() string {
	if s.Secure {
		return HTTPSProtocol
	}
	return HTTPProtocol
}

func Load() (*App, error) {
	var cfg App
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
