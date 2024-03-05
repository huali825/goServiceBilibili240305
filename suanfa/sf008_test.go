package suanfa

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestSuanfa008(t *testing.T) {
	go func() {
		defer fmt.Println("A.defer")
		go func() {
			defer fmt.Println("B.defer")
			go func() {
				defer fmt.Println("C.defer")
				runtime.Goexit()
				fmt.Println("C func")
			}()
			fmt.Println("B func")
		}()
		fmt.Println("A func")
	}()

	time.Sleep(time.Hour)
}
