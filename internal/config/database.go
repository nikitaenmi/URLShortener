package config

type Database struct {
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	User     string `env:"DB_USER"`
	DBName   string `env:"DB_DBNAME"`
	Password string `env:"DB_PASSWORD"`
	SSLMode  string `env:"DB_SSLMODE"`
}
