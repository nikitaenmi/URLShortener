package redirect

import (
	"net/http"

	"github.com/nikitaenmi/URLShortener/internal/database"
	"github.com/nikitaenmi/URLShortener/internal/database/models"
)

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	shortCode := r.URL.Path[1:] // Извлекаем код из URL

	var link models.Link
	result := database.Migration().Where("short_code = ?", shortCode).First(&link)
	if result.Error != nil {
		http.Error(w, "Ссылка не найдена", http.StatusNotFound)
		return
	}

	// Перенаправляем на оригинальный URL
	http.Redirect(w, r, link.OriginalURL, http.StatusMovedPermanently)
}
