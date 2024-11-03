package ctx_pkg

import (
	"context"
	"fmt"
	"time"
)

func printNumbers(ctx context.Context) {

	for i := 1; i <= 10; i++ {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println(i)
			time.Sleep(1 * time.Second)
		}

	}
}
func Task1() {
	fmt.Println("Task-1    *********************")
	ctx, cancel := context.WithCancel(context.Background())
	go printNumbers(ctx)
	time.Sleep(3 * time.Second)
	fmt.Println("(cancellation)")
	cancel()
}
