package crud_intface

import (
	"fmt"
	"session-20/cn_to_dbase"
	"session-20/dec_mdls"
)

func createNewBook() {
	newBook := dec_mdls.Book{Author: "John Doe", Title: "Go Programming", PublishedYear: 2023}
	if cn_to_dbase.DB.Create(&newBook).Error == nil {
		fmt.Println("Book record inserted successfully!")
	}
}

func printAllBooks() {
	var books []dec_mdls.Book
	result := cn_to_dbase.DB.Find(&books)
	if result.Error == nil {
		fmt.Println("Books in database:")
		for _, book := range books {
			fmt.Printf("ID: %d, Title: %s, Author: %s, Year: %d\n", book.ID, book.Title, book.Author, book.PublishedYear)
		}

	}

}
func Task3() {
	createNewBook()
	printAllBooks()
}
