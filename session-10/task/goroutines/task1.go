package goroutines

import (
	"fmt"
	"time"
)

var nums = []int{1, 2, 3, 4, 5}

func Task1() {
	fmt.Println("Main function started")

	go func() {
		for _, num := range nums {
			fmt.Printf("%d\n", num)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	time.Sleep(time.Second)
	fmt.Println("Main function ended")
}
