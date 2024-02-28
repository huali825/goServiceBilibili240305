package channel

import (
	"testing"
	"time"
)

func TestLoopChannel(t *testing.T) {
	ch := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
			time.Sleep(time.Millisecond * 10)
		}
		close(ch)
	}()

	for val := range ch {
		t.Log(val)
	}
}
