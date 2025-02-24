package shorten

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nikitaenmi/URLShortener/internal/database"
	"github.com/nikitaenmi/URLShortener/internal/database/models"
	"github.com/teris-io/shortid"
)

func ShortenURL(w http.ResponseWriter, r *http.Request) {

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
	shortCode, err := shortid.Generate()
	if err != nil {
		http.Error(w, "Ошибка генерации кода", http.StatusInternalServerError)
		return
	}

	// Сохранение в базу данных
	link := models.Link{
		OriginalURL: request.URL,
		ShortCode:   shortCode,
	}
	if err := database.Migration().Create(&link).Error; err != nil {
		http.Error(w, "Ошибка сохранения в базу данных", http.StatusInternalServerError)
		return
	}

	fmt.Println(shortCode)

	// Возвращаем короткую ссылку
	response := map[string]string{
		"short_url": "http://localhost:8080/" + shortCode,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
