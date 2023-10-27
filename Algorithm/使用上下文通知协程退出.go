package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	// 启动协程 A
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		// 执行协程 A 的任务
		fmt.Println("协程 A 正在执行")
		time.Sleep(2 * time.Second)
		cancel() // 通知协程 B 可以退出
	}(ctx)

	// 启动协程 B
	go func(ctx context.Context) {
		// 在协程 B 中等待协程 A 完成或接收退出通知
		select {
		case <-ctx.Done():
			fmt.Println("协程 B 收到退出通知，正在退出")
		}
	}(ctx)

	// 等待协程 A 完成
	wg.Wait()
	fmt.Println("协程 A 已完成")
}
