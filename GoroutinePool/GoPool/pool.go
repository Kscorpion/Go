package GoPool

import (
	"fmt"
)

type Task struct {
	taskId int
	f      func() error
}

type Pool struct {
	workerNum  int
	EntryChan  chan *Task
	workerChan chan *Task
}

func NewTask(id int, f func() error) *Task {
	return &Task{
		taskId: id,
		f:      f,
	}
}

func (t *Task) execute() {
	t.f()
}

func NewPool(num int) *Pool {
	return &Pool{
		workerNum:  num,
		EntryChan:  make(chan *Task),
		workerChan: make(chan *Task),
	}
}

//每个worker分别从workerChan中进行获取任务
func (p *Pool) worker(id int) {
	for task := range p.workerChan {
		task.execute()
		fmt.Println("workerId:", id, "taskId:", task.taskId, "is done")
	}
	close(p.workerChan)
}

func (p *Pool) Run() {
	//首先根据客户端定制的worker数进行创建worker协程
	for i := 0; i < p.workerNum; i++ {
		go p.worker(i)
	}
	//从外部写入的任务管道中获取任务写入worker协程管道
	for task := range p.EntryChan {
		p.workerChan <- task
	}
}
