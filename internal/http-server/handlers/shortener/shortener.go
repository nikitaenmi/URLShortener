package shortener

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/nikitaenmi/URLShortener/internal/config"
	"github.com/nikitaenmi/URLShortener/internal/database"
	"github.com/nikitaenmi/URLShortener/internal/database/models"
)

// Интерфейс для генерации алиасов
type Generater interface {
	Generate() (string, error)
}

type Logger interface {
	Info(msg string, args ...any)
	Error(msg string, args ...any)
}

func ShortenerURL(w http.ResponseWriter, r *http.Request, cfg config.App, gen Generater, log Logger) {
	var request struct {
		URL string `json:"url"`
	}

	// Декодируем JSON-запрос
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		log.Error("Invalid request", err)
		return
	}

	// Генерация алиаса для URL
	alias, err := gen.Generate()
	if err != nil {
		http.Error(w, "Error generating alias", http.StatusInternalServerError)
		log.Error("Error generating alias: %v\n", err)
		return
	}

	// Сохранение в базу данных
	link := models.Link{
		OriginalURL: request.URL,
		Alias:       alias,
	}

	// Создаинтерфейсем подключение к базе данных
	db := database.Migration(cfg.Database)

	// Используем подключение для выполнения операций
	if err := db.Create(&link).Error; err != nil {
		log.Error("Failed to create link:", err)
		return
	}

	log.Info("Link created successfully!", alias)

	// Возвращаем короткую ссылку
	response := map[string]string{
		"short_url": fmt.Sprintf("http://%s:%s/%s", cfg.Server.Host, cfg.Server.Port, alias),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
