package suanfa

import (
	"fmt"
	"sync"
	"testing"
)

func printFruit(fruit string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		ch <- fruit
	}
}

func TestSuanfa002(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan string)

	wg.Add(3)
	go printFruit("apple", ch, &wg)
	go printFruit("banana", ch, &wg)
	go printFruit("icecream", ch, &wg)

	go func() {
		for i := 0; i < 12; i++ {
			fmt.Println(<-ch)
		}
	}()

	wg.Wait()
	close(ch)
}
