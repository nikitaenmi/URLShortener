package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/nikitaenmi/URLShortener/internal/database"
	"github.com/nikitaenmi/URLShortener/internal/http-server/handlers/redirect"
	"github.com/nikitaenmi/URLShortener/internal/http-server/handlers/shorten"
)

const localhost string = ":8080"

func main() {
	db := database.Migration()

	repo := &redirect.UrlDB{DB: db}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	http.HandleFunc("/shorten", shorten.ShortenURL)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		redirect.RedirectURL(w, r, repo, logger)
	})

	slog.Info("Сервер Запущен")
	err := http.ListenAndServe(localhost, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
