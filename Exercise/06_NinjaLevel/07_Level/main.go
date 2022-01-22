package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x := func(numb ...int) []int {
		var res []int
		for _, v := range numb {
			if v%2 == 0 {
				res = append(res, v)
			}
		}
		return res
	}
	y := x(numbers...)
	fmt.Println(y)
}
