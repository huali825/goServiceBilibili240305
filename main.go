package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	go func() {
		defer fmt.Println("A.defer")
		go func() {
			defer fmt.Println("B.defer")
			runtime.Goexit()
			//go func() {
			//	defer fmt.Println("C.defer")
			//	runtime.Goexit()
			//	fmt.Println("C func")
			//}()
			fmt.Println("B func")
		}()
		fmt.Println("A func")
	}()

	time.Sleep(time.Hour)
}
