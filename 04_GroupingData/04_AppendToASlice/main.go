package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	x := []int{4, 5, 7, 8, 45}
	fmt.Println(x)
	x = append(x, rand.Intn(100))
	fmt.Println(x)
	fmt.Println(x[len(x) - 1])
	y := []int{234, 456, 678, 987}
	x = append(x, y...)
	fmt.Println(x)
	z := []int{51, 144, 1412, 5141, 51244}
	x = append(x, z[:3]...)
	fmt.Println(x)
}
