package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 45
	}()

	fmt.Println(<-c)
}
