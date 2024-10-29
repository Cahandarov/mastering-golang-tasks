package sync_atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count int64

func increment(wg *sync.WaitGroup) {
	atomic.AddInt64(&count, 10)
	defer wg.Done()
}

func Task9() {
	fmt.Println("Task-9   **********************")

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&wg)
	}
	wg.Wait()
	fmt.Println("Final atomic counter value:", count)
}
