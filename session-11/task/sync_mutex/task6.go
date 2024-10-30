package sync_mutex

import (
	"fmt"
	"sync"
)

var studentGrades = make(map[string]int)

func updateGrades(name string, score int, mux *sync.Mutex, wg *sync.WaitGroup) {
	mux.Lock()
	studentGrades[name] = score
	mux.Unlock()
	defer wg.Done()
}
func Task6() {
	fmt.Println("Task-6   **********************")
	mux := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(4)
	go updateGrades("Ali", 70, &mux, &wg)
	go updateGrades("Ali", 100, &mux, &wg)
	go updateGrades("Hasan", 80, &mux, &wg)
	go updateGrades("Aylin", 100, &mux, &wg)
	wg.Wait()
	fmt.Println("Final Grades:", studentGrades)
}
