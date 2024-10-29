package sync_waitgroup

import (
	"fmt"
	"sync"
	"time"
)

func Task3() {
	fmt.Println("Task-3   **********************")
	wg := sync.WaitGroup{}

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func() {
			fmt.Printf("Goroutine %d starting\n", i)
			time.Sleep(1 * time.Second)
			fmt.Printf("Goroutine %d finished\n", i)
			defer wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("All goroutines have completed")
}
