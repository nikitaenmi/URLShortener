package domain

type Url struct {
	ID          int `gorm:"primarykey"`
	OriginalURL string
	Alias       string
}

type URLFilter struct {
	Alias string
}

type UrlRepo interface {
	Create(url Url) error
	URLFind(params URLFilter) (*Url, error)
}