package main

import (
	"e-library/books"
	"e-library/users"
	"fmt"
)

func main() {
	fmt.Println("  _____      _     _ _                          ")
	fmt.Println(" | ____|    | |   (_) |__  _ __ __ _ _ __ _   _ ")
	fmt.Println(" |  _| _____| |   | | '_ \\| '__/ _` | '__| | | |")
	fmt.Println(" | |__|_____| |___| | |_) | | | (_| | |  | |_| |")
	fmt.Println(" |_____|    |_____|_|_.__/|_|  \\__,_|_|   \\__, |")
	fmt.Println("                                          |___/ ")

	books.AvailableBooks()
	fmt.Println()
	users.BorrowedBooks()
	fmt.Println()
	fmt.Println("End of Report.")

}
