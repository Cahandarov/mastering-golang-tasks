package rate_limit_and_throttling

import (
	"fmt"
	"sync"
	"time"
)

func job(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	data := <-ch
	fmt.Println(data)
	time.Sleep(1 * time.Second)
}

func Task4() {
	fmt.Println("Task-4    *********************")
	jobLimit := 5
	rateLimit := 2
	ticker := time.NewTicker(1 * time.Second)
	wg := sync.WaitGroup{}
	wg.Add(jobLimit)
	defer ticker.Stop()

	ch := make(chan string, rateLimit)

	for i := 0; i < jobLimit; i++ {
		ch <- "Job started"
		go job(ch, &wg)

		<-ticker.C
	}

	wg.Wait()
	fmt.Println("(program ends)")
}
