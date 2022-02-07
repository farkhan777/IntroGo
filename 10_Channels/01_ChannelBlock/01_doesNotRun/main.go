package main

import "fmt"

func main() {
	c := make(chan int)

	c <- 45

	fmt.Println(<-c)
}
