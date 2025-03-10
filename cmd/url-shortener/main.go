package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"
	"github.com/nikitaenmi/URLShortener/internal/config"
	"github.com/nikitaenmi/URLShortener/internal/database"
	"github.com/nikitaenmi/URLShortener/internal/http-server/handlers/redirect"
	"github.com/nikitaenmi/URLShortener/internal/http-server/handlers/shortener"
)

func main() {
	var cfg config.App
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(".env not found")
	}

	db := database.Migration(cfg.Database)

	repo := &redirect.UrlDB{DB: db}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	http.HandleFunc("/shortener", shortener.ShortenerURL)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		redirect.RedirectURL(w, r, repo, logger)
	})

	slog.Info("Сервер запущен")
	err = http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port), nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
