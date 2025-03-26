package domain


type Link struct {
	ID          int `gorm:"primarykey"`
	OriginalURL string
	Alias       string
}

