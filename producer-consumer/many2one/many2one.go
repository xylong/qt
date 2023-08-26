package many2one

import (
	"qt/producer-consumer/out"
	"sync"
)

const (
	taskNum = 10000
	num     = 100
)

var (
	wg1         = &sync.WaitGroup{}
	wg2         = &sync.WaitGroup{}
	taskChannel = make(chan *Task, 10)
)

type Task struct {
	ID int
}

func (t *Task) Run() {
	out.Println(t.ID)
}

func producer(c chan<- *Task, start, n int) {
	for i := start; i < start+n; i++ {
		c <- &Task{
			ID: i,
		}
	}
}

func consumer(c <-chan *Task) {
	for task := range c {
		task.Run()
	}
}

func Exec() {
	for i := 0; i < taskNum; i += num {
		wg1.Add(1)
		wg2.Add(1)
		go func(i int) {
			defer wg1.Done()
			defer wg2.Done()
			producer(taskChannel, i, num)
		}(i)
	}

	wg1.Add(1)
	go func() {
		defer wg1.Done()
		consumer(taskChannel)
	}()

	wg2.Wait()
	go close(taskChannel)

	wg1.Wait()
}
