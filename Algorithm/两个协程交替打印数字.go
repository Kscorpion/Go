package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	ctx, cancelFunc := context.WithTimeout(context.TODO(), 5*time.Second)
	go func(ctx context.Context, cancelFunc1 context.CancelFunc, ch chan int) {
		for i := 1; i <= 100; i += 1 {
			<-ch
			if i%2 == 0 {
				fmt.Println(i)
				if i == 100 {
					close(ch)
					cancelFunc1()
				}
			}
		}
	}(ctx, cancelFunc, ch)

	go func(ctx context.Context) {
		for i := 1; i <= 100; i += 1 {
			ch <- 0
			if i%2 == 1 {
				fmt.Println(i)
			}
		}
	}(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("结束输出")
	}
}
