package main

import (
	"qt/producer-consumer/many2one"
	"qt/producer-consumer/out"
)

func main() {
	o := out.NewOut()
	go o.Output()

	//out.Println(1)
	//out.Println(2)
	//out.Println(3)

	//one2one.Exec()
	//one2many.Exec()
	many2one.Exec()

	// 等待退出信号
	//s := make(chan os.Signal)
	//signal.Notify(s, syscall.SIGINT, syscall.SIGALRM)
	//<-s
}
