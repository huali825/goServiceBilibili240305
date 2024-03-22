package suanfa

import (
	"fmt"
	"sync"
	"testing"
)

func TestSuanfa011(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)

	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})

	go printLetter02("a", ch1, ch2, &wg)
	go printLetter02("b", ch2, ch3, &wg)
	go printLetter02("c", ch3, ch1, &wg)

	// 启动第一个协程
	ch1 <- struct{}{}

	// 等待所有协程完成
	wg.Wait()

	// 关闭管道
	close(ch1)
	close(ch2)
	close(ch3)
}

// 使用三个管道实现三个协程同步顺序打印a b c
func printLetter02(letter string, prevCh, nextCh chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		// 等待上一个协程通知
		<-prevCh
		fmt.Print(letter)
		// 发送信号给下一个协程
		nextCh <- struct{}{}
	}
}
