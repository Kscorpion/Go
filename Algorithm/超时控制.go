package main

import (
	"context"
	"fmt"
	"time"
)

var ()

func main() {
	c := make(chan int)
	ctx, cancelfunc := context.WithTimeout(context.TODO(), 5*time.Second)
	list := []int{23, 32, 78, 43, 76, 65, 345, 762}
	go a(ctx, list, 345, c)
	go a(ctx, list, 345, c)
	select {
	case <-c:
		cancelfunc()
		return
	case <-ctx.Done():
		fmt.Println("超时了")
		cancelfunc()
		return
	}
}

func a(cancel context.Context, list []int, flg int, c chan<- int) {
	for _, val := range list {
		if val == flg {
			//time.Sleep(10 * time.Second)
			fmt.Println("Found it")
			c <- 10
			close(c)
			break
		}
	}
}
