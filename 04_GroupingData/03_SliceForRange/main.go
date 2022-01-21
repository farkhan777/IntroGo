package main

import "fmt"

func main()  {
	x := []int{4, 5, 7, 8, 45}
	fmt.Println(len(x))
	fmt.Print(x[:])
	fmt.Println()
	fmt.Print(x[3:])
	fmt.Println()
	fmt.Print(x[3:4])
	fmt.Println()

	for i, v := range x {
		fmt.Println(i, v)
	}

	for i := 0; i < len(x); i++ {
		fmt.Print(x[i], " ")
	}
}
