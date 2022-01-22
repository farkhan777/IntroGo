package main

import "fmt"

func main() {
	fmt.Println(foo(10))
	age, name := bar()
	fmt.Println(age, name)
}

func foo(n int) int {
	return n * n
}

func bar() (int, string) {
	return 21, "Farkhan Hamzah Firdaus"
}
