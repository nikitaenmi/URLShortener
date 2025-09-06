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
	"github.com/nikitaenmi/URLShortener/internal/http-server/handlers"
	"github.com/nikitaenmi/URLShortener/internal/repository"
	"github.com/nikitaenmi/URLShortener/internal/services"
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

	repo := repository.NewUrl(db)
	svc := services.NewUrl(repo)
	handler := handlers.NewUrl(svc, logger, cfg.Server)

	http.HandleFunc("/shortener", handler.ShortenerURL)
	http.HandleFunc("/", handler.RedirectURL)

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	logger.Info("Server is running", slog.String("address", srv.Addr))
	err = srv.ListenAndServe()
	if err != nil {
		logger.Error("Server not running", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
