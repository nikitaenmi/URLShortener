package redirect

import (
	"log/slog"
	"net/http"

	"github.com/nikitaenmi/URLShortener/internal/database/models"
	"gorm.io/gorm"
)

type UrlDB struct {
	DB *gorm.DB
}

func (r *UrlDB) FinderOriginalCode(generatedCode string) (*models.Link, error) { // по короткому коду, который был сгенерированный, возращает оригинальную ссылку
	var link models.Link

	result := r.DB.Where("generated_code = ?", generatedCode).First(&link)
	if result.Error != nil {
		return nil, result.Error
	}

	return &link, nil
}

type Finder interface {
	FinderOriginalCode(shortCode string) (*models.Link, error)
}

type Logger interface {
	Info(msg string, args ...any)
	Error(msg string, args ...any)
}

func RedirectURL(w http.ResponseWriter, r *http.Request, repo Finder, log Logger) {
	shortID := r.URL.Path[1:] // Извлекаем сгенерированный код из URL, когда пользовател отправил предоставленную короткую ссылку

	link, err := repo.FinderOriginalCode(shortID)
	if err != nil {
		log.Error("Ссылка не найдена", slog.String("shortID", shortID), slog.Any("error", err))
		http.Error(w, "Ссылка не найдена", http.StatusNotFound)
		return
	}

	log.Info("Перенаправление", slog.String("shortID", shortID), slog.String("original_url", link.OriginalURL))
	http.Redirect(w, r, link.OriginalURL, http.StatusMovedPermanently)
}
