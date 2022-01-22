package main

import (
	"fmt"
)

func main() {
	score := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	a := sum(score...)
	fmt.Println("All numbers:", a)

	t := evenSum(sum, score...)
	fmt.Println("Even numbers:", t)

	d := oddSum(sum, score...)
	fmt.Println("Odd numbers:", d)
}

func sum(x ...int) int {
	y := 0
	for _, v := range x {
		y += v
	}

	return y
}

func evenSum(r func(x ...int) int, p ...int) int {
	var res []int
	for _, v := range p {
		if v%2 == 0 {
			res = append(res, v)
		}
	}

	total := r(res...)
	return total
}

func oddSum(r func(x ...int) int, p ...int) int {
	var res []int
	for _, v := range p {
		if v%2 != 0 {
			res = append(res, v)
		}
	}

	return r(res...)
}
