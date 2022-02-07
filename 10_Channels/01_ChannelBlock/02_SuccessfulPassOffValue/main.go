package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 45
		// you can use only one
		c <- 47
	}()

	fmt.Println(<-c)
	fmt.Println(<-c)
}
