package httpserver

import (
	"github.com/nikitaenmi/URLShortener/internal/domain"
)

func (r *UrlRequest) ToDomain() domain.Url {
	return domain.Url{
		OriginalURL: r.OriginalURL,
	}
}

func ToUrlItemResponse(alias string, host, port string) UrlItemResponse {
	return UrlItemResponse{
		ShortURL: "http://" + host + ":" + port + "/" + alias,
	}
}

func (r *UrlListRequest) ValidateAndSetDefaults() {
	if r.Page < 1 {
		r.Page = 1
	}
	if r.Limit < 1 || r.Limit > 50 {
		r.Limit = 10
	}
}

func (r *UrlListRequest) ToDomain() domain.Paginator {
	return domain.Paginator{
		Page:  r.Page,
		Limit: r.Limit,
	}
}

func ToUrlListResponse(urlList *domain.UrlList) UrlListResponse {
	return UrlListResponse{
		URLs:  urlList.Items,
		Total: int(urlList.Total),
	}
}
