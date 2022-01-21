package main

import "fmt"

var x int

type hotdog int

var y hotdog

func main() {
	x = 45
	y = 47
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	x = int(y)
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
