package shorten

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	GeneratedCode, err := shortid.Generate()
	if err != nil {
		http.Error(w, "Ошибка генерации кода", http.StatusInternalServerError)
		return
	}
	// TO DO: architecture

	// // Сохранение в базу данных
	// link := models.Link{
	// 	OriginalURL:   request.URL,
	// 	GeneratedCode: GeneratedCode,
	// }

	// if err := database.Migration(cfg.Database).Create(&link).Error; err != nil {
	// 	http.Error(w, "Ошибка сохранения в базу данных", http.StatusInternalServerError)
	// 	return
	// }

	// fmt.Println(GeneratedCode)

	// Возвращаем короткую ссылку
	response := map[string]string{
		"short_url": "http://localhost:8080/" + GeneratedCode,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
