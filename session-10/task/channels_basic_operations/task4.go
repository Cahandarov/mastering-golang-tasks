package channels_basic_operations

import (
	"fmt"
)

var numSlc = []int{1, 2, 3, 4, 5}

func Task4() {
	fmt.Println("Task4  *****************************")

	ch := make(chan int)
	go func() {
		for _, n := range numSlc {
			ch <- n
		}
		close(ch)

	}()
	for c := range ch {
		fmt.Println(c)
	}
	fmt.Println("Channel closed")
}
