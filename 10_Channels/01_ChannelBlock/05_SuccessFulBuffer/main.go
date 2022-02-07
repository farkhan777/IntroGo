package main

import "fmt"

func main() {
	c := make(chan int, 2)

	c <- 45
	c <- 47

	fmt.Println(<-c)
	fmt.Println(<-c)
}
