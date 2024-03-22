package suanfa

import (
	"fmt"
	"testing"
	"time"
)

func TestSuanfa012(t *testing.T) {
	msg := make(chan int)
	msg2 := make(chan int)
	msg3 := make(chan int)

	go printFirst(msg, msg3)
	go printSecond(msg, msg2)
	go printThird(msg2, msg3)

	time.Sleep(time.Second * 2)
}

var POOL = 100

// 为什么每一个goroutine都要从从1运行到100呢？
// 因为每个数字都要传递到3个协程当中，如果有跳跃式的遍历，那么会导致后面的协程缺少数字
func printFirst(p chan int, r chan int) {
	for i := 1; i <= POOL; i++ {
		p <- i
		if i%3 == 1 {
			fmt.Println("goroutine-1:", i)
		}
		<-r
	}
}

func printSecond(p chan int, q chan int) {
	for i := 1; i <= POOL; i++ {
		<-p
		if i%3 == 2 {
			fmt.Println("goroutine-2:", i)
		}
		q <- i
	}
}

func printThird(q chan int, r chan int) {
	for i := 1; i <= POOL; i++ {
		<-q
		if i%3 == 0 {
			fmt.Println("goroutine-3:", i)
		}
		r <- i
	}
}
