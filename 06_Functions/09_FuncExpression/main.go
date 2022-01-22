package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("My first func expression")
	}
	f()

	g := func(x int) {
		fmt.Println("I was born in:", x)
	}
	g(2000)
}
