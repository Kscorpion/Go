package main

import (
	"context"
	"fmt"
	"time"
)

type API interface {
	Call(ctx context.Context, id string)
}

type MyAPI struct {
}

func (a *MyAPI) Call(ctx context.Context, id string) {
	select {
	case <-ctx.Done():
		fmt.Println("这个协程退出了奥", id) // 返回上下文错误，表示取消原因
	}
}

func main() {
	ctx1, cancel1 := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel1()
	ctx2, cancel2 := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel2()
	ctx3, cancel3 := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel3()
	api := &MyAPI{}

	go api.Call(ctx1, "1")
	go api.Call(ctx2, "2")
	go api.Call(ctx3, "3")

	time.Sleep(10 * time.Second)
}
