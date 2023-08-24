package one2one

import (
	"qt/producer-consumer/out"
	"sync"
)

type Task struct {
	ID int64
}

func (t *Task) Run() {
	out.Println(t.ID)
}

var taskChannel = make(chan *Task, 10)

const taskNum = 10000

// 生产者
func producer(c chan<- *Task) {
	var i int64

	for i = 1; i < taskNum; i++ {
		c <- &Task{ID: i}
	}

	close(c) // 生产者不关闭channel，消费者不会退出循环一直阻塞，wg一直等不到协程结束就会产生死锁
}

// 消费者
func consumer(c <-chan *Task) {
	// 生产者端关闭channel才会中断循环
	for task := range c {
		task.Run()
	}
}

func Exec() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		producer(taskChannel)
	}()

	go func() {
		defer wg.Done()
		consumer(taskChannel)
	}()

	wg.Wait()
	out.Println("执行完毕")
}
