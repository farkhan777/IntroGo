package main

import "fmt"

func main() {
	c := make(chan int)

	// Send
	go foo(c)

	// Receive
	bar(c)

	fmt.Printf("About to exit")
}

// Send
func foo(c chan<- int) {
	c <- 97
}

// Receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}
