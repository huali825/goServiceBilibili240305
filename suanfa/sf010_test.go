package suanfa

import (
	"fmt"
	"sync"
	"testing"
)

func TestSuanFa010(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)

	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})

	go printerNum(1, "a", ch1, ch2, &wg)
	go printerNum(2, "b", ch2, ch3, &wg)
	go printerNum(3, "c", ch3, ch1, &wg)
	ch1 <- struct{}{}

	wg.Wait()

}

func printerNum(
	num int,
	str string,
	a chan struct{},
	b chan struct{},
	wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 34; i++ {
		<-a
		fmt.Println(str, ":", i*3+num)
		b <- struct{}{}
	}

	if str == "a" {
		<-a
	}

}
