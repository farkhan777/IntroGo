package main

import "fmt"

func main() {
	cr := make(<-chan int)

	go func() {
		cr <- 100
	}()

	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("c\t%T\n", cr)
}
