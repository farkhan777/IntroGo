package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 100
	}()

	fmt.Println(<-c)

	fmt.Printf("------\n")
	fmt.Printf("c\t%T\n", c)
}
