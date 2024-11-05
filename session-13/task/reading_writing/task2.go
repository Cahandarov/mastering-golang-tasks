package reading_writing

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var lines int

func createFile() {
	file, err := os.OpenFile("task/reading_writing/story.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println("Error occurred:", err)
	}
	defer file.Close()

	_, err = file.WriteString("Once upon a time, there was a brave knight.\nHe fought dragons and saved villages.\nThe people loved him dearly.\n")
	if err != nil {
		fmt.Println("Error occurred on writing story:", err)
	}
}

func readFile() {
	file, err := os.Open("task/reading_writing/story.txt")
	if err != nil {
		fmt.Println("Error occurred:", err)
	}
	defer file.Close()

	r := bufio.NewReader(file)

	for {
		_, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error occurred while reading:", err)
			return
		} else {
			lines++
		}
	}
	fmt.Printf("Total lines in file: %d\n", lines)
}

func Task2() {
	fmt.Println("Task-2   ********************")
	createFile()
	readFile()
}
