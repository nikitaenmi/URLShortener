package httpserver

import (
	"github.com/nikitaenmi/URLShortener/internal/constants"
	"github.com/nikitaenmi/URLShortener/internal/domain"
)

func (r *URLRequest) ToDomain() domain.URL {
	return domain.URL{
		OriginalURL: r.OriginalURL,
	}
}

func ToURLItemResponse(protocol, host, port, alias string) URLItemResponse {
	return URLItemResponse{
		ShortURL: protocol + "://" + host + ":" + port + "/" + alias,
	}
}

func (r *Paginator) ValidateAndSetDefaults() {
	if r.Page < 1 {
		r.Page = constants.DefaultPage
	}
	if r.Limit < 1 || r.Limit > constants.MaxLimit {
		r.Limit = constants.DefaultLimit
	}
}

func (r *Paginator) ToDomain() domain.Paginator {
	return domain.Paginator{
		Page:  r.Page,
		Limit: r.Limit,
	}
}

func ToURLListResponse(urlList *domain.URLList) URLListResponse {
	return URLListResponse{
		URLs:  urlList.Items,
		Total: int(urlList.Total),
	}
}
