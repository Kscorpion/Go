package main

import (
	"context"
	"fmt"
	"time"
)

type API interface {
	Call() (string, error)
}

type MyAPI struct {
}

func (a *MyAPI) Call0() (string, error) {
	time.Sleep(1 * time.Second)
	return "success call0", nil
}
func (a *MyAPI) Call1() (string, error) {
	time.Sleep(3 * time.Second)
	return "success call1", nil
}

func (a *MyAPI) Call2() (string, error) {
	time.Sleep(5 * time.Second)
	return "success call2", nil
}

func main() {
	ctx, calcel := context.WithTimeout(context.Background(), 2*time.Second)
	defer calcel()
	api := &MyAPI{}
	result := make(chan string)

	go func() {
		data, _ := api.Call0()
		result <- data
	}()

	go func() {
		data, _ := api.Call1()
		result <- data
	}()

	go func() {
		data, _ := api.Call2()
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
