package shorten

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"
	"github.com/nikitaenmi/URLShortener/internal/config"
	"github.com/nikitaenmi/URLShortener/internal/database"
	"github.com/nikitaenmi/URLShortener/internal/database/models"
	"github.com/teris-io/shortid"
)



func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var cfg config.App
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(".env not found")
	}

	var request struct {
		URL string `json:"url"`
	}

	// Декодируем JSON-запрос
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		fmt.Println(err)
		fmt.Println(http.StatusBadRequest)
		return
	}

	// Генерация короткого кода
	GeneratedCode, err := shortid.Generate()
	if err != nil {
		http.Error(w, "Ошибка генерации кода", http.StatusInternalServerError)
		return
	}
	// TO DO: architecture

	// Сохранение в базу данных
	link := models.Link{
		OriginalURL:   request.URL,
		GeneratedCode: GeneratedCode,
	}

	// Создаем подключение к базе данных
	db := database.Migration(cfg.Database)

	// Используем подключение для выполнения операций
	if err := db.Create(&link).Error; err != nil {
		fmt.Println("Failed to create link:", err)
		return
	}

	fmt.Println("Link created successfully!")

	fmt.Println(GeneratedCode)

	// Возвращаем короткую ссылку
	response := map[string]string{
		"short_url": "http://localhost:8080/" + GeneratedCode,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
