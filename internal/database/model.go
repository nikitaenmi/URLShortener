package database

type URL struct {
	ID          int `gorm:"primaryKey"`
	OriginalURL string
	Alias       string `gorm:"uniqueIndex"`
}
