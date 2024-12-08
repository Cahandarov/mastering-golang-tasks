package db_sql_pkg

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitializeDB() {
	if DB == nil {
		if err := connectDB(); err != nil {
			fmt.Println(err)
			return
		}
	}
}

func CloseDB() {
	if DB != nil {
		DB.Close()
		DB = nil
	}
}

func connectDB() error {
	connStr := "host=localhost port=5432 user=student password=securepass dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	DB = db

	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("database ping issue: %w", err)
	}

	if err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}
	fmt.Println("Connection successful")
	return nil
}

func Task3() {
	fmt.Println("Task 3  ****************")

	InitializeDB()

	CloseDB()

}
