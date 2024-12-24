package model

type Book struct {
	ID            uint
	Title         string  `gorm:"not null"`
	Author        string  `gorm:"not null"`
	Price         float64 `gorm:"not null"`
	PublishedYear int
}
