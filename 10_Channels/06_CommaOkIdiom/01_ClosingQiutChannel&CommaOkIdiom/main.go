package main

import (
	"fmt"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(even, odd, quit)

	receive(even, odd, quit)

	fmt.Println("About to exit")
}

// Receive
func receive(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("The value received from the even channel:", v)
		case v := <-o:
			fmt.Println("The value received from the odd channel:", v)
		case i, ok := <-q:
			if !ok {
				fmt.Println("From comma ok bit", i)
				return
			} else {
				fmt.Println("Form comma ok bit", i)
			}
		}
	}
}

// Send
func send(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)
}
