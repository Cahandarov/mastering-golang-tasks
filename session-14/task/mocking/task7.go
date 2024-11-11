package mocking

import (
	"errors"
)

type Database interface {
	GetUserByID(id int) (User, error)
}

type User struct {
	Id   int
	Name string
}

var ErrorUserNotFound = errors.New("failed to find user")

func GetUser(db Database, id int) (User, error) {
	user, err := db.GetUserByID(id)
	if err != nil {
		return User{}, ErrorUserNotFound
	}
	return user, nil
}

// func Task7() {
// 	fmt.Println("Task-7    **************8")

// }
