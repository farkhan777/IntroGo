package main

import "fmt"

func main() {
	fmt.Println(incrementer()())
	fmt.Println(incrementer()())
	fmt.Println(incrementer()())

	a := incrementer()
	b := incrementer()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())

	{
		y := 45
		fmt.Println(y)
	}
	//fmt.Println(y)
}

func incrementer() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
