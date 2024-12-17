package cn_to_dbase

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func openDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=testdb port=5432 sslmode=disable TimeZone=Asia/Baku"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	DB = db

	fmt.Println("Connected to PostgreSQL database successfully!")

	DB = DB.Debug()
}

func Task1() {
	openDB()
}
