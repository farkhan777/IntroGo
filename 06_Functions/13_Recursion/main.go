package main

import "fmt"

func main() {
	fmt.Println(factorial(10))

	fmt.Println(foo(10))

	fmt.Println(loopFact(10))
}

func factorial(numb int) int {
	if numb == 0 {
		return 1
	}
	return numb * factorial(numb-1)
}

func foo(numb int) int {
	res := 1
	for i := numb; i > 0; i-- {
		res *= numb
		numb--
	}

	return res
}

func loopFact(numb int) int {
	res := 1
	for ; numb > 0; numb-- {
		res *= numb
	}

	return res
}
