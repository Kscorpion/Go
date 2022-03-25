package main

import (
	"GoPool/GoPool"
	"fmt"
	"time"
)

func task() error {
	fmt.Println(time.Now(), "Do something")
	return nil
}

func main() {
	p := GoPool.NewPool(3)
	id := 0
	go func() {
		for {
			p.EntryChan <- GoPool.NewTask(id, task)
			id++
		}
		close(p.EntryChan)
	}()
	p.Run()
}
