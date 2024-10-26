package buffered_unbuffered_channels

import (
	"fmt"
	"time"
)

var numbs = []int{10, 20, 30}

func Task5() {
	fmt.Println("Task5  *****************************")

	ch := make(chan int, len(numbs))
	for _, n := range numbs {
		ch <- n
	}

	fmt.Println("Sent values into buffered channel")
	close(ch)

	go func() {
		for c := range ch {
			fmt.Println(c)
		}
		fmt.Println("All values received")

	}()
	time.Sleep(100 * time.Millisecond)
}
