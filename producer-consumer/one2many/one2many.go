package one2many

import (
	"qt/producer-consumer/out"
	"sync"
)

const taskNum = 10000

var taskChaneel = make(chan *Task)

type Task struct {
	ID int
}

func (t *Task) Run() {
	out.Println(t.ID)
}

func producer(c chan<- *Task) {
	for i := 1; i <= taskNum; i++ {
		c <- &Task{
			ID: i,
		}
	}

	close(c)
}

func consumer(c <-chan *Task) {
	for task := range c {
		task.Run()
	}
}

func Exec() {
	wg := &sync.WaitGroup{}

	// 1个生产者
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		producer(taskChaneel)
	}(wg)

	// n个消费者
	for i := 0; i < taskNum; i++ {
		if i%100 == 0 {
			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				consumer(taskChaneel)
			}(wg)
		}
	}

	wg.Wait()
	out.Println("执行完毕")
}
