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
	"github.com/nikitaenmi/URLShortener/internal/repository"
)

func main() {
	var cfg config.App
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(".env not found")
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	db, err := database.Connect(cfg.Database)
	if err != nil {
		log.Fatal("Database init", err.Error())
	}

	repo := &repository.UrlDB{DB: db}

	http.HandleFunc("/shortener", func(w http.ResponseWriter, r *http.Request) {
		shortener.ShortenerURL(w, r, repo, cfg.Server, logger)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		redirect.RedirectURL(w, r, repo, logger)
	})

	logger.Info("Server is running")
	err = http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port), nil)
	if err != nil {
		logger.Error("Server not running", err)
		os.Exit(1)
	}
}
