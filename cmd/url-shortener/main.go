package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/nikitaenmi/URLShortener/internal/config"
	"github.com/nikitaenmi/URLShortener/internal/database"
	"github.com/nikitaenmi/URLShortener/internal/http-server/handlers"
	"github.com/nikitaenmi/URLShortener/internal/http-server/middleware"
	"github.com/nikitaenmi/URLShortener/internal/lib/generator"
	"github.com/nikitaenmi/URLShortener/internal/lib/logger"
	"github.com/nikitaenmi/URLShortener/internal/repository"
	"github.com/nikitaenmi/URLShortener/internal/services"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})) // Логгер для логов в мейне (slog)

	cfg, err := config.Load()
	if err != nil {
		log.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	db, err := database.Connect(cfg.Database)
	if err != nil {
		log.Error("database init failed", "error", err.Error())
		os.Exit(1)
	}

	repo := repository.NewURL(db)
	aliasGenerator, err := generator.New(cfg.Generator)
	if err != nil {
		log.Error("generator init failed", "error", err.Error())
		os.Exit(1)
	}
	svc := services.NewURL(repo, aliasGenerator)

	loggerWithContext := logger.NewLoggerWithContex()
	e := echo.New()
	e.Use(middleware.RequestIDMiddleware())
	e.Use(logger.EchoRequestLogger(loggerWithContext))
	e.HTTPErrorHandler = middleware.HTTPErrorHandler

	handler := handlers.NewURL(svc, loggerWithContext, cfg)

	e.POST("/urls", handler.Create)
	e.GET("/urls/:id", handler.Read)
	e.PUT("/urls/:id", handler.Update)
	e.DELETE("/urls/:id", handler.Delete)
	e.GET("/urls", handler.List)

	e.GET("/:alias", handler.Redirect)

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	log.Info("server is running", "address:", srv.Addr)

	if err := e.StartServer(srv); err != nil {
		log.Error("server not running", "error", err.Error())
		os.Exit(1)
	}
}
