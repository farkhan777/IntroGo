package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("This is foo.")
}

func bar() {
	fmt.Println("This is bar.")
}
