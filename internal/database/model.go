package database

type Url struct {
	ID          int `gorm:"primaryKey"`
	OriginalURL string
	Alias       string `gorm:"uniqueIndex"`
}
