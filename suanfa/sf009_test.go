package suanfa

import (
	"fmt"
	"testing"
)

func TestSuanfa009(t *testing.T) {
	c := make(chan int, 2)

	go func() {
		defer fmt.Println("goroutine 结束")
		fmt.Println("goroutine 运行中")

		c <- 1001
	}()

	num := <-c
	fmt.Println("num的值为", num)
}
