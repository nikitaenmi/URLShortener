package config

import "time"

type Server struct {
	Host         string        `env:"SERVER_HOST"`
	Port         string        `env:"SERVER_PORT"`
	TimeoutRead  time.Duration `env:"SERVER_TIMEOUT_R_S"`
	TimeoutWrite time.Duration `env:"SERVER_TIMEOUT_W_S"`
	IdleTimeout  time.Duration `env:"SERVER_IDLETIMEOUT_S"`
}
