package main

import "fmt"

func main()  {
	// COMPOSITE LITERAL
	x := []int{1,7,8,34,83,145,35}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[len(x) - 1])
}

// a SLICE allows you to group together VALUES of the same TYPE
