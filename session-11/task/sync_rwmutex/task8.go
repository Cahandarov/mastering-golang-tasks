package sync_rwmutex

import (
	"fmt"
	"sync"
)

func increment(count *int, mut *sync.RWMutex, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		mut.Lock()
		*count++
		fmt.Println("Writer updated counter:", *count)
		mut.Unlock()
	}()
	go func() {
		defer wg.Done()
		mut.RLock()
		fmt.Println("Reader accessed counter:", *count)
		mut.RUnlock()
	}()
}
func Task8() {
	fmt.Println("Task-8   **********************")

	count := 0
	mut := sync.RWMutex{}
	wg := sync.WaitGroup{}
	for i := 0; i < 4; i++ {
		wg.Add(2)
		increment(&count, &mut, &wg)
	}
	wg.Wait()
}
