package main

import "fmt"

func main() {
	c := make(chan int, 1)

	c <- 45
	c <- 47

	fmt.Println(<-c)
}
