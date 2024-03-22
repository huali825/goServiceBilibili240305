package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var maxNum int = 100
	wg.Add(3)
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)
	go printNum(&maxNum, ch1, ch2, &wg)
	go printNum(&maxNum, ch2, ch3, &wg)
	go printNum(&maxNum, ch3, ch1, &wg)

	ch1 <- 1

	wg.Wait()

}

func printNum(maxNum *int, a chan int, b chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		ptrNum := <-a
		if ptrNum == *maxNum+1 {
			b <- ptrNum
			return
		}

		fmt.Println("goroutine:", ptrNum)
		ptrNum++
		b <- ptrNum
	}
}
func printNum2(maxNum *int, a chan int, b chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		ptrNum := <-a
		if ptrNum > *maxNum {
			b <- ptrNum
			return
		}

		fmt.Println("goroutine:", ptrNum)
		ptrNum++
		b <- ptrNum
	}
}
