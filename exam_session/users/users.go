package users

import (
	"e-library/books"
	"errors"
	"fmt"
)

const maxBorrowLimit = 3

var ErrorBorrowLimit = errors.New("borrow limit reached")
var ErrorReturnWrongBook = errors.New("this book have not borrowed")
var ErrorAlreadyBorrowedBook = errors.New("this book is already borrowed")

type User struct {
	UserID        int
	Name          string
	BorrowedBooks []int
}

var users = map[int]User{
	1: {UserID: 1, Name: "Ali"},
	2: {UserID: 2, Name: "Babek"},
	3: {UserID: 3, Name: "Cabir", BorrowedBooks: []int{1, 2}},
	4: {UserID: 4, Name: "Diana", BorrowedBooks: []int{6, 9, 10}},
	5: {UserID: 5, Name: "Elvira"},
}

func (u User) CheckBorrowedBooks() string {
	if len(u.BorrowedBooks) > 0 {
		return fmt.Sprintf("%s borrowed %d books\n", u.Name, len(u.BorrowedBooks))
	} else {
		return fmt.Sprintf("%s hasn't borrowed any books yet\n", u.Name)
	}
}

func canBorrowMore(userID int) bool {
	if val, ok := users[userID]; ok {
		if len(val.BorrowedBooks) < maxBorrowLimit {
			return true
		}
	}
	return false
}

func checkAlreadyBorrowed(userID, bookID int) error {
	if val, ok := users[userID]; ok {
		for _, book := range val.BorrowedBooks {
			if book == bookID {
				return ErrorAlreadyBorrowedBook
			}
		}
	}
	return nil
}

func borrowBook(userID, bookID int) error {
	status, errAvailable := books.BookAvailable(bookID) //checks status of book, if invalid id panics
	if errAvailable != nil {
		return errAvailable
	}

	if status && canBorrowMore(userID) { //if limit not reached allows to borrow
		errBorrowedSame := checkAlreadyBorrowed(userID, bookID) //checks borrowed same book already
		if errBorrowedSame != nil {
			return errBorrowedSame
		}

		user := users[userID]
		user.BorrowedBooks = append(users[userID].BorrowedBooks, bookID)
		users[userID] = user

		for i, book := range books.Books {
			if book.BookID == bookID {
				books.Books[i].Status = "borrowed"
			}
		}
		return nil //return nil if borrowed successfully

	} else {
		return ErrorBorrowLimit
	}
}

func returnBook(userID, bookID int) error {
	user := users[userID]
	for i, book := range user.BorrowedBooks {
		if book == bookID {

			user.BorrowedBooks = append(user.BorrowedBooks[:i], user.BorrowedBooks[i+1:]...)
			users[userID] = user
			for i, book := range books.Books {
				if book.BookID == bookID {
					books.Books[i].Status = "available"
				}
			}

			return nil
		}
	}
	return ErrorReturnWrongBook
}

func printBorrowedBooksByUsers() {
	count := 0
	for _, user := range users {
		if len(user.BorrowedBooks) > 0 {
			fmt.Println("2. Borrowed Books by Users:")
			break
		}
	}
	for _, user := range users {
		if len(user.BorrowedBooks) > 0 {
			count++
			fmt.Printf("- User %d (%s):\n", count, user.Name)
		}
		for _, b := range user.BorrowedBooks {
			for _, book := range books.Books {
				if book.BookID == b {
					fmt.Printf("\t- \"%s\" by %s\n", book.Title, book.Author)
				}
			}

		}
	}
}

func BorrowedBooks() {
	printBorrowedBooksByUsers()

	//borrowBook(1, 2)
	////fmt.Println("borrow same book:", borrowBook(1, 2))
	//borrowBook(1, 4)
	////borrowBook(1, 4)
	//borrowBook(1, 41)
	////borrowBook(1, 3)
	//borrowBook(1, 9)
	////fmt.Println("try to borrow unavailable book:", borrowBook(1, 6))
	////fmt.Println("try to reach borrow limit", borrowBook(1, 7))
	//returnBook(1, 3)
	////fmt.Println("try to return wrong book:", returnBook(1, 13))
	//borrowBook(1, 5)
}
