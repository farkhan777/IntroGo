package main

import (
	"fmt"
	"runtime"
)

func main() {
	x := 10
	y := 10
	c := foo(x, y)

	for i := 0; i < x*y; i++ {
		fmt.Println(i, <-c)
	}

	fmt.Println("GOROUTINES", runtime.NumGoroutine())
}

func foo(x, y int) <-chan int {
	c := make(chan int)

	for i := 0; i < x; i++ {
		go func() {
			for j := 0; j < y; j++ {
				c <- j
			}
		}()
		fmt.Println("GOROUTINES", runtime.NumGoroutine())
	}

	return c
}
