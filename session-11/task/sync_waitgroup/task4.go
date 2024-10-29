package sync_waitgroup

import (
	"fmt"
	"sync"
)

var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8}

func calculateSum(slice []int, ch chan int, wg *sync.WaitGroup) {

	defer wg.Done()
	var sum int = 0
	for i := 0; i < len(slice); i++ {
		sum += slice[i]
	}
	ch <- sum
}
func Task4() {
	fmt.Println("Task-4   **********************")
	wg := sync.WaitGroup{}
	ch := make(chan int, 2)
	wg.Add(2)

	go calculateSum(numbers[:4], ch, &wg)
	partialSum1 := <-ch
	go calculateSum(numbers[4:], ch, &wg)
	partialSum2 := <-ch

	fmt.Println("Partial sum 1:", partialSum1)
	fmt.Println("Partial sum 2:", partialSum2)
	fmt.Println("Total sum:", partialSum1+partialSum2)

	wg.Wait()
	close(ch)
}
