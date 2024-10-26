package goroutines

import (
	"fmt"
	"time"
)

var numbers = []int{1, 2, 3, 4, 5}
var letters = []string{"A", "B", "C", "D", "E"}

func Task2() {
	fmt.Println("Task2  *****************************")

	go func() {
		for _, num := range numbers {
			fmt.Printf("%d\n", num)
			time.Sleep(300 * time.Millisecond)
		}
	}()

	go func() {
		for _, l := range letters {
			fmt.Printf("%s\n", l)
			time.Sleep(200 * time.Millisecond)
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Main finished")
}
