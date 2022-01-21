package main

import "fmt"

const (
	a int = iota * 10
	b int = iota * 10
	c float64 = iota * 10
)

func main()  {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
