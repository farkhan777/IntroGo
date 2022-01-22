package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	x := foo(numbers...)
	fmt.Println(x)

	y := bar(numbers)
	fmt.Println(y)
}

func foo(n ...int) int {
	res := 0
	for _, v := range n {
		res += v
	}

	return res
}

func bar(n []int) int {
	res := 0
	for _, v := range n {
		res += v
	}

	return res
}
