package sync_mutex

import (
	"fmt"
	"sync"
)

func increment(count *int, mut *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mut.Lock()
	*count++
	mut.Unlock()
}
func Task5() {
	fmt.Println("Task-5   **********************")

	count := 0
	mut := sync.Mutex{}
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go increment(&count, &mut, &wg)
	}
	wg.Wait()
	fmt.Println("Final counter value:", count)
}
