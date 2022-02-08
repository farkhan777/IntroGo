package main

import "fmt"

func main() {
	c := make(chan int, 1)
	cr := make(<-chan int, 1) // receive
	cs := make(chan<- int, 1) // send

	fmt.Println("-----")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)
}
