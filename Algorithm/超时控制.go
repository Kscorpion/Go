package main

import (
	"context"
	"fmt"
	"time"
)

type API interface {
	Call(time.Duration) (string, error)
}

type MyAPI struct {
}

func (a *MyAPI) Call(t time.Duration) (string, error) {
	time.Sleep(t * time.Second)
	return "success call", nil
}

func main() {
	ctx, calcel := context.WithTimeout(context.Background(), 2*time.Second)
	defer calcel()
	api := &MyAPI{}
	result := make(chan string)

	go func() {
		data, _ := api.Call(1)
		result <- data
	}()

	go func() {
		data, _ := api.Call(3)
		result <- data
	}()

	go func() {
		data, _ := api.Call(5)
		result <- data
	}()

	for i := 0; i < 3; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("timeout")
		case res := <-result:
			fmt.Println(res)
		}
	}
}
