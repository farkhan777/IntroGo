package main

import "fmt"

func main() {
	func(n int) {
		fmt.Println("I am", n, "years old")
	}(21)
}
