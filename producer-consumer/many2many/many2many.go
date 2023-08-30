package many2many

import (
	"qt/producer-consumer/out"
	"time"
)

type Task struct {
	ID int
}

func (t *Task) Run() {
	out.Println(t.ID)
}

var (
	taskChannel = make(chan *Task, 10)
	quit        = make(chan struct{})
)

const (
	taskNum = 10000
)

func producer(t chan<- *Task, quit chan struct{}) {
	var i int

	for {
		if i >= taskNum {
			i = 0
		}
		i++

		select {
		case t <- &Task{ID: i}:
		case <-quit:
			out.Println("生产者退出")
			return
		}
	}
}

func consumer(t <-chan *Task, quit chan struct{}) {
	for {
		select {
		case task := <-t:
			if task.ID != 0 {
				task.Run()
			}
		case <-quit:
			// 收到完成通知时channel不一定消费完了
			for task := range t {
				if task.ID != 0 {
					task.Run()
				}
			}
			out.Println("消费者退出")
			return
		}
	}
}

func Exec() {
	go producer(taskChannel, quit)
	go producer(taskChannel, quit)
	go producer(taskChannel, quit)
	go producer(taskChannel, quit)
	go producer(taskChannel, quit)
	go producer(taskChannel, quit)

	go consumer(taskChannel, quit)
	go consumer(taskChannel, quit)

	time.Sleep(time.Second * 2)
	close(quit)
	close(taskChannel)
	time.Sleep(time.Second * 2)

	out.Println(len(taskChannel))
}
