package wraperr

import (
	"errors"
	"fmt"
	"os"
)

func openFile(filename string) error {
	var _, err = os.Open(filename)
	if err != nil {
		return fmt.Errorf("file not found: %w", err)
	} else {
		return nil
	}

}
func Task3() {
	fmt.Println("Task 3 - *****************************")
	err := openFile("nonexistent.txt")
	fmt.Println("Wrapped Error: ", err)
	var UnwrappedError = errors.Unwrap(err)
	fmt.Println("Original Error: ", UnwrappedError)
}
