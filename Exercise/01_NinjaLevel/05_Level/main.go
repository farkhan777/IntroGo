package main

import "fmt"

type age int
var x age
var y int

func main()  {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 21
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
