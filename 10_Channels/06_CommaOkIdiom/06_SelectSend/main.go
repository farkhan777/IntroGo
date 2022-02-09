package main

import "fmt"

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	foo(c, quit)
}

func foo(c, q chan int) {
	x := 3
	for {
		select {
		case c <- x:
			x += x
		case <-q:
			fmt.Println("About to exit")
			return
		}
	}
}
