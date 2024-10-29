package select_statement

import (
	"fmt"
	"time"
)

func Task2() {
	fmt.Println("Task-2   **********************")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(4 * time.Second)
		ch1 <- "Hello from ch1"
	}()

	go func() {
		time.Sleep(4 * time.Second)
		ch2 <- "Hello from ch2"
	}()

	select {
	case data := <-ch1:
		fmt.Println("Data received:", data)
	case data := <-ch2:
		fmt.Println("Data received:", data)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout occurred: no message received")
	}
}
