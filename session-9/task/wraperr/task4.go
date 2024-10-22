package wraperr

import (
	"fmt"
	"os"
)

func openFile2(filename string) error {
	var file, errOpen = os.Open(filename)
	if errOpen != nil {
		return fmt.Errorf("failed to open file: %w", errOpen)
	}

	if errRead := readFile(file); errRead != nil {
		return fmt.Errorf("failed to read file: %w", errOpen)
	}

	return nil
}
func readFile(f *os.File) error {
	return fmt.Errorf("file reading issue")
}
func Task4() {
	fmt.Println("Task 4 - *****************************")
	err := openFile2("nonexistent.txt")
	fmt.Println("Error: ", err)
}
