package buffered_unbuffered_channels

import "fmt"

func Task6() {
	fmt.Println("Task6  *****************************")

	chStr := make(chan string)

	chStr <- "Hello"
}
