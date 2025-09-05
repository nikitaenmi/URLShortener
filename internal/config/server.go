package config

import "time"

type Server struct {
	Host        string        `env:"SERVER_HOST"`
	Port        string        `env:"SERVER_PORT"`
	Timeout     time.Duration `env:"SERVER_TIMEOUT_S"`
	IdleTimeout time.Duration `env:"SERVER_IDLETIMEOUT_S"`
}
