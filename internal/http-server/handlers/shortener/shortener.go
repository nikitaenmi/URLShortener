package shortener

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/nikitaenmi/URLShortener/internal/config"
	"github.com/nikitaenmi/URLShortener/internal/database"
	"github.com/nikitaenmi/URLShortener/internal/database/models"
	"github.com/teris-io/shortid"
	"gorm.io/gorm"
)

type UrlDB struct {
	DB *gorm.DB
}

func (r *UrlDB) Create(link *models.Link) error {
	return r.DB.Create(link).Error
}

func ShortenerURL(w http.ResponseWriter, r *http.Request, cfg config.App) {
	var request struct {
		URL string `json:"url"`
	}

	// Декодируем JSON-запрос
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		fmt.Println(err)
		fmt.Println(http.StatusBadRequest)
		return
	}

	// Генерация алиаса для URL
	aliace, err := shortid.Generate()
	if err != nil {
		http.Error(w, "Error generating alias", http.StatusInternalServerError)
		return
	}

	// Сохранение в базу данных
	link := models.Link{
		OriginalURL: request.URL,
		Aliace:      aliace,
	}

	// Создаем подключение к базе данных
	db := database.Migration(cfg.Database)

	// Используем подключение для выполнения операций
	if err := db.Create(&link).Error; err != nil {
		fmt.Println("Failed to create link:", err)
		return
	}

	fmt.Println("Link created successfully!")
	fmt.Println(aliace)

	// Возвращаем короткую ссылку
	response := map[string]string{
		"short_url": fmt.Sprintf("http://%s:%s/%s", cfg.Server.Host, cfg.Server.Port, aliace),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
