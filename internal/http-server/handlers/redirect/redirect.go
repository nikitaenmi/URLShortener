package redirect

import (
	"log/slog"
	"net/http"

	"github.com/nikitaenmi/URLShortener/internal/database/models"
)

type Finder interface {
	FinderOriginalCode(shortCode string) (*models.Link, error)
}

type Logger interface {
	Info(msg string, args ...any)
	Error(msg string, args ...any)
}

func RedirectURL(w http.ResponseWriter, r *http.Request, repo Finder, log Logger) {
	aliace := r.URL.Path[1:] // Извлекаем алиас из URL, когда пользовател отправил предоставленную короткую ссылку

	link, err := repo.FinderOriginalCode(aliace)
	if err != nil {
		log.Error("Link not found", slog.String("aliace", aliace), slog.Any("error", err))
		http.Error(w, "Link not found", http.StatusNotFound)
		return
	}

	log.Info("Redirection", slog.String("aliace", aliace), slog.String("originalURL", link.OriginalURL))
	http.Redirect(w, r, link.OriginalURL, http.StatusMovedPermanently)
}
