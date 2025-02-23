package main

import (
	"fmt"
	"net/http"
	db "urlShortener/internal/database"
	"urlShortener/internal/http-server/handlers/redirect"
	"urlShortener/internal/http-server/handlers/shorten"
)

func main() {
	db.Migration()

	http.HandleFunc("/shorten", shorten.ShortenURL)
	http.HandleFunc("/", redirect.RedirectURL)

	fmt.Println("Сервер запущен")

	http.ListenAndServe(":8080", nil)

}
