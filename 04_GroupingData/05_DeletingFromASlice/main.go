package main

import "fmt"

func main()  {
	x := []int{4, 5, 7, 8, 45}
	fmt.Println(x)
	y := []int{234, 456, 678, 987}
	fmt.Println(y)
	x = append(x, y...)
	fmt.Println(x)
	x = append(x[:5], x[7:]...)
	fmt.Println(x)
}
