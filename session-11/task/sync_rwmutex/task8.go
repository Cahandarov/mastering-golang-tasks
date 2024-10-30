package sync_rwmutex

import (
	"fmt"
	"sync"
)

var counter int

func readCounter(rwMux *sync.RWMutex, wg *sync.WaitGroup) {
	rwMux.RLock()
	fmt.Println("Reader accessed counter:", counter)
	rwMux.RUnlock()
	defer wg.Done()
}
func writeCounter(rwMux *sync.RWMutex, wg *sync.WaitGroup) {
	rwMux.Lock()
	counter++
	rwMux.Unlock()
	fmt.Println("Writer updated counter:", counter)
	defer wg.Done()
}
func Task8() {
	fmt.Println("Task-8   **********************")

	mux := sync.RWMutex{}
	wg := sync.WaitGroup{}
	wg.Add(4)
	go readCounter(&mux, &wg)
	go writeCounter(&mux, &wg)
	go readCounter(&mux, &wg)
	go writeCounter(&mux, &wg)
	wg.Wait()
}
