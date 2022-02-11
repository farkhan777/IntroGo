package main

import "fmt"

func main() {
	defer foo()
	bar()
	goo()
}

func foo() {
	fmt.Println("This is foo.")
}

func bar() {
	fmt.Println("This is bar.")
}

func goo() {
	fmt.Println("This is goo.")
}
