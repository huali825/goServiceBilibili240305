package channel

import (
	"fmt"
	"testing"
)

func TestLoopChannel(t *testing.T) {
	//声明
	var ch1 chan int
	var ch11 chan struct{} //空结构体 用来做信号

	// 普通定义,  不带容量的
	ch2 := make(chan int)
	//带容量的定义
	ch3 := make(chan int, 2)

	go func() {
		for i := 0; i < 20; i++ {
			//向ch2中 发 数据
			ch2 <- i
			//time.Sleep(time.Millisecond * 10)
		}
		close(ch2)
	}()

	for val := range ch2 {
		//t.Log(val)
		fmt.Print(val, " ")
	}

	fmt.Println(ch1, ch11, ch2, ch3, "打印完毕")
}
