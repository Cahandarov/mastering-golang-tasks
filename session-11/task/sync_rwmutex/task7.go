package sync_rwmutex

import (
	"fmt"
	"sync"
)

var users = make(map[string]int)

func updateUsers(name string, age int, mux *sync.RWMutex, wg *sync.WaitGroup) {
	wg.Add(2)
	go func() {
		mux.Lock()
		users[name] = age
		mux.Unlock()
		defer wg.Done()
	}()
	go func() {
		mux.RLock()
		//fmt.Println(users[name])
		mux.RUnlock()
		defer wg.Done()
	}()
}
func Task7() {
	fmt.Println("Task-7   **********************")
	mux := sync.RWMutex{}
	wg := sync.WaitGroup{}

	updateUsers("Ali", 25, &mux, &wg)
	updateUsers("Ali", 45, &mux, &wg)
	updateUsers("Nursultan", 54, &mux, &wg)
	updateUsers("Hatice", 18, &mux, &wg)

	wg.Wait()
	fmt.Println("User data:", users)
}
