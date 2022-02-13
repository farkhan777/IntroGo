package main

import "fmt"

func main() {
	x, err := fmt.Println("Hello, World!")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(x)
}
