package structs

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func Task1() {

	book := Book{
		Title:  "The Go Programming Language",
		Author: "Alan A. A. Donovan",
		Pages:  380,
	}
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\n", book.Title, book.Author, book.Pages)

}
