package main

import "fmt"

func main() {
	f := make(chan int)
	b := make(chan int)

	go func() {
		for i := 0; i < 5000; i++ {
			f <- i
		}
	}()

	go func() {
		for i := 0; i < 5000; i++ {
			b <- i
		}
	}()
	foo(f, b)
}

func foo(f, b chan int) {
	for i := 0; i < 10000; i++ {
		select {
		case v := <-f:
			fmt.Println("From foo:", v)
		case v := <-b:
			fmt.Println("From bar:", v)
		}
	}
	fmt.Println("About to exit")
}
