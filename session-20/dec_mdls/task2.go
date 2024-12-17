package dec_mdls

import (
	"fmt"
	"session-20/cn_to_dbase"
	"time"
)

type Book struct {
	ID            uint   `gorm:"primaryKey"`
	Title         string `gorm:"size:255;not null"`
	Author        string `gorm:"size:100"`
	PublishedYear int
	CreatedAt     time.Time
}

func createTable() {
	err := cn_to_dbase.DB.AutoMigrate(&Book{})
	if err != nil {
		fmt.Printf("Failed to migrate table: %v", err)
	}
}
func Task2() {
	createTable()
}
