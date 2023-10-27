package main

import (
	"context"
	"fmt"
	"time"
)

type API interface {
	Call() (string, error)
}

type MyAPI struct{}

func (a *MyAPI) Call() (string, error) {
	// 模拟一个耗时的操作
	time.Sleep(3 * time.Second)
	return "API调用成功", nil
}

func main() {
	// 创建一个带有超时控制的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	api := &MyAPI{}
	result := make(chan string)

	// 启动一个协程来执行接口调用
	go func() {
		data, err := api.Call()
		if err != nil {
			result <- err.Error()
			return
		}
		result <- data
	}()

	select {
	case <-ctx.Done():
		fmt.Println("接口调用超时")
	case res := <-result:
		fmt.Println(res)
	}
}
