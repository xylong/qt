package main

import (
	"os"
	"os/signal"
	"qt/producer-consumer/many2many"
	"qt/producer-consumer/out"
	"syscall"
)

func main() {
	o := out.NewOut()
	go o.Output()

	//out.Println(1)
	//out.Println(2)
	//out.Println(3)

	//one2one.Exec()
	//one2many.Exec()
	//many2one.Exec()
	many2many.Exec()

	// 等待退出信号
	s := make(chan os.Signal)
	signal.Notify(s, syscall.SIGINT, syscall.SIGALRM)
	<-s
}
