package sync_rwmutex

import (
	"fmt"
	"sync"
)

var users = make(map[string]int)

func readFromMap(name string, mux *sync.RWMutex, wg *sync.WaitGroup) {
	mux.RLock()
	_ = users[name]
	mux.RUnlock()
	defer wg.Done()
}
func writeToMap(name string, grade int, mux *sync.RWMutex, wg *sync.WaitGroup) {
	mux.Lock()
	users[name] = grade
	mux.Unlock()
	defer wg.Done()
}
func Task7() {
	fmt.Println("Task-7   **********************")
	mux := sync.RWMutex{}
	wg := sync.WaitGroup{}

	wg.Add(6)
	go writeToMap("Garay", 20, &mux, &wg)
	go writeToMap("Ali", 25, &mux, &wg)
	go writeToMap("Medina", 28, &mux, &wg)
	go readFromMap("Garay", &mux, &wg)
	go readFromMap("Ali", &mux, &wg)
	go readFromMap("Medina", &mux, &wg)

	wg.Wait()
	fmt.Println("User data:", users)
}
