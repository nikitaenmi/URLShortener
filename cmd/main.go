package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/nikitaenmi/URLShortener/internal/database"
	"github.com/nikitaenmi/URLShortener/internal/http-server/handlers/redirect"
	"github.com/nikitaenmi/URLShortener/internal/http-server/handlers/shorten"
)

func main() {
	database.Migration()

	http.HandleFunc("/shorten", shorten.ShortenURL)
	http.HandleFunc("/", redirect.RedirectURL)

	fmt.Println("Сервер запущен")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
