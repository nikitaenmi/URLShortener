package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/nikitaenmi/URLShortener/internal/config"
	"github.com/nikitaenmi/URLShortener/internal/database"
	"github.com/nikitaenmi/URLShortener/internal/http-server/handlers"
	"github.com/nikitaenmi/URLShortener/internal/http-server/middleware"
	"github.com/nikitaenmi/URLShortener/internal/lib/generator"
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

	aliasGenerator, err := generator.New(cfg.Generator)
	if err != nil {
		log.Fatal("Generatocr init", err.Error())
	}

	svc := services.NewUrl(repo, aliasGenerator)
	handler := handlers.NewUrl(svc, logger, cfg.Server)

	e := echo.New()
	e.Use(middleware.RequestIDMiddleware())
	e.Use(middleware.ErrorHandler(logger))

	// CRUDL - OPERATIONS
	e.POST("/urls", handler.Create)
	e.GET("/urls/:id", handler.Read)
	e.PUT("/urls/:id", handler.Update)
	e.DELETE("/urls/:id", handler.Delete)
	e.GET("/urls", handler.List)

	e.GET("/r/:alias", handler.Redirect)

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	logger.Info("Server is running", slog.String("address", srv.Addr))

	if err := e.StartServer(srv); err != nil {
		logger.Error("Server not running", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
