package redirect

import (
	"net/http"

	"github.com/nikitaenmi/URLShortener/internal/database/models"
)

type Finder interface {
	FinderAlias(alias string) (*models.Link, error)
}

type Logger interface {
	Info(msg string, args ...any)
	Error(msg string, args ...any)
}

func RedirectURL(w http.ResponseWriter, r *http.Request, repo Finder, log Logger) {
	alias := r.URL.Path[1:]

	link, err := repo.FinderAlias(alias)
	if err != nil {
		log.Error("Link not found")
		http.Error(w, "Link not found", http.StatusNotFound)
		return
	}

	log.Info("Redirection")
	http.Redirect(w, r, link.OriginalURL, http.StatusMovedPermanently)
}
