package out

import "fmt"

type Out struct {
	data chan interface{}
}

var out *Out

func NewOut() *Out {
	if out == nil {
		out = &Out{
			data: make(chan interface{}),
		}
	}

	return out
}

func (o *Out) Output() {
	for {
		select {
		case i := <-o.data:
			fmt.Println(i)
		}
		
		fmt.Println("out put") // 命中一次case才执行一次for
	}
}

func Println(i interface{}) {
	out.data <- i
}
