package main

import "fmt"

func main() {
	defer foo()
	bar()
	goo()

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
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
