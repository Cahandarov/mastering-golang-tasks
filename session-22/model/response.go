package model

type BookResponse struct {
	Title         string  `gorm:"not null"`
	Author        string  `gorm:"not null"`
	Price         float64 `gorm:"not null"`
	PublishedYear int
}

type AuthResponse struct {
	Token string `json:"accessToken"`
}
