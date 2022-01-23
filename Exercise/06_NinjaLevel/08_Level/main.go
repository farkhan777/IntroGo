package main

import "fmt"

func main() {
	y := foo(100)
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())

	fmt.Println()

	r := foo(100)
	fmt.Println(r())
	fmt.Println(r())
	fmt.Println(r())

	fmt.Println()

	fmt.Println(foo(1000)())
}

func foo(n int) func() int {
	x := n
	return func() int {
		x--
		return x
	}
}
