package config

type Server struct {
	Host        string `env:"SERVER_HOST"`
	Port        string `env:"SERVER_PORT"`
	Timeout     int    `env:"SERVER_TIMEOUT"`
	IdleTimeout int    `env:"SERVER_IDLETIMEOUT"`
}
