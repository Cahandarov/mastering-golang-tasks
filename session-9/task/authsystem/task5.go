package authsystem

import (
	"errors"
	"fmt"
)

type ErrorResponse struct {
	Message string
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("%s", e.Message)
}

var (
	UserNotFoundError = ErrorResponse{
		Message: "user not found",
	}
	InvalidPasswordError = ErrorResponse{
		Message: "invalid password",
	}
)

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

func authenticate(username, password string) (string, error) {
	if passValue, ok := users[username]; ok {
		if passValue == password {
			return "", nil
		}
		return username, InvalidPasswordError
	}
	return username, UserNotFoundError
}

func Task5() {
	fmt.Println("Task 5 - *****************************")

	user, aut := authenticate("user5", "password1")
	if aut == nil {
		fmt.Println("Login successful!")
	} else if errors.Is(aut, InvalidPasswordError) {
		fmt.Printf("Error: %s for user: %s\n", aut, user)
	} else if errors.Is(aut, UserNotFoundError) {
		fmt.Printf("Error: %s: %s\n", aut, user)
	}
}
