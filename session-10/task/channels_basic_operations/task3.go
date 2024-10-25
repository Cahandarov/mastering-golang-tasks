package channels_basic_operations

import (
	"fmt"
	"time"
)

func Task3() {
	fmt.Println("Task3  *****************************")
	c := make(chan int)
	go func() {
		time.Sleep(500 * time.Millisecond)
		c <- 42
	}()
	data := <-c
	fmt.Printf("Received value: %d\n", data)
}
