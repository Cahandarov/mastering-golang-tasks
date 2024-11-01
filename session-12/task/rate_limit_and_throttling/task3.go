package rate_limit_and_throttling

import (
	"fmt"
	"sync"
	"time"
)

func throttling(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Task executed")
}
func Task3() {
	fmt.Println("Task-3    *********************")
	wg := sync.WaitGroup{}
	wg.Add(5)
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for i := 0; i < 5; i++ {
		<-ticker.C
		go throttling(&wg)
	}
	wg.Wait()
	fmt.Println("(ticker stopped)")
}
