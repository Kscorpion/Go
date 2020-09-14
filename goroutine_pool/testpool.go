package main

import (
	"fmt"
	"time"
	"pool"
)

func task() error {
	fmt.Println(time.Now(), "Do something")
	return nil
}

func main() {
	p := pool.NewPool(3)
	id := 0
	go func() {
		for {
			p.EntryChan <- pool.NewTask(id, task)
			id++
		}
	}()
	p.Run()
}
