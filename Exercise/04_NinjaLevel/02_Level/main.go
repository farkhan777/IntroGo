package main

import "fmt"

func main()  {
	arr := []int{31, 41, 15, 63, 36, 35 ,354, 562, 525, 363}

	for _, v := range arr {
		fmt.Printf("%v ", v)
	}

	fmt.Printf("\n%T", arr)
}
