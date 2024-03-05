package suanfa

import (
	"fmt"
	"testing"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		//t.Log("i的值为:", i)
		fmt.Println("i的值为:", i)
		time.Sleep(time.Second)
	}
}

func TestSuanfa007(t *testing.T) {
	go newTask()
	i := 0
	for {
		i++
		//t.Log("i的值为:", i)
		fmt.Println("i的值为:", i)
		time.Sleep(time.Second)
	}
}
