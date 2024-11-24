package books

import (
	"errors"
	"fmt"
)

var ErrorUnavailableBook = errors.New("book not available in e-library")
var ErrorOnAddingNewBook = errors.New("something went wrong on adding new book")

type Book struct {
	BookID int
	Title  string
	Author string
	Status string
}

var Books = []Book{
	{BookID: 1, Title: "The Go Programming Guide", Author: "John Doe", Status: "borrowed"},
	{BookID: 2, Title: "Clean Code", Author: "Robert Martin", Status: "borrowed"},
	{BookID: 3, Title: "Concurrency in Practice", Author: "Jane Smith", Status: "available"},
	{BookID: 4, Title: "Microservices Design", Author: "Alan Turing", Status: "available"},
	{BookID: 5, Title: "Data Structures", Author: "Grace Hopper", Status: "available"},
	{BookID: 6, Title: "The Art of Computer Prog", Author: "Donald Knuth", Status: "borrowed"},
	{BookID: 7, Title: "Algorithms Unlocked", Author: "Thomas Cormen", Status: "available"},
	{BookID: 8, Title: "Software Architecture", Author: "Martin Fowler", Status: "available"},
	{BookID: 9, Title: "System Design", Author: "Scott Meyers", Status: "borrowed"},
	{BookID: 10, Title: "Effective Go", Author: "Rob Pike", Status: "borrowed"},

	{BookID: 11, Title: "The Go Programming Guide", Author: "John Doe", Status: "available"},
	{BookID: 12, Title: "Clean Code", Author: "Robert Martin", Status: "available"},
	{BookID: 13, Title: "Concurrency in Practice", Author: "Jane Smith", Status: "available"},
	{BookID: 14, Title: "Microservices Design", Author: "Alan Turing", Status: "available"},
	{BookID: 15, Title: "Data Structures", Author: "Grace Hopper", Status: "available"},
	{BookID: 16, Title: "The Art of Computer Prog", Author: "Donald Knuth", Status: "available"},
	{BookID: 17, Title: "Algorithms Unlocked", Author: "Thomas Cormen", Status: "available"},
	{BookID: 18, Title: "Software Architecture", Author: "Martin Fowler", Status: "available"},
	{BookID: 19, Title: "System Design", Author: "Scott Meyers", Status: "available"},
	{BookID: 20, Title: "Effective Go", Author: "Rob Pike", Status: "available"},
}

func AddBook(id int, title string, author string) error {
	if id < 0 || title == "" || author == "" {
		return ErrorOnAddingNewBook
	}
	newBook := Book{BookID: id, Title: title, Author: author, Status: "available"}
	Books = append(Books, newBook)
	return nil
}

func BookAvailable(bookId int) (bool, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic", err)
		}
	}()

	for _, book := range Books {
		if book.BookID == bookId {
			if book.Status == "available" {
				return true, nil
			} else {
				return false, ErrorUnavailableBook
			}
		}
	}
	panic("invalid book ID entered")
}

func AvailableBooks() {
	if len(Books) > 0 {
		fmt.Println("1. Available Books:")
		for _, book := range Books {
			if book.Status == "available" {
				fmt.Printf("- %s by %s\n", book.Title, book.Author)
			}
		}
	}
}
