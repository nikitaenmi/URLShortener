package models

type Link struct {
	ID            int `gorm:"primarykey"`
	OriginalURL   string
	GeneratedCode string
}
