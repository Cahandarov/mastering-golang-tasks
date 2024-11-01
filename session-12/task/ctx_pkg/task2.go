package ctx_pkg

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func lateResponse(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Task cancelled due to timeout")
			return
		case <-time.After(6 * time.Second):
			fmt.Println("Late response")
			return
		}
	}
}

func Task2() {
	fmt.Println("Task-2    *********************")
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	wg.Add(1)
	go lateResponse(ctx, &wg)

	defer cancel()
	wg.Wait()
}
