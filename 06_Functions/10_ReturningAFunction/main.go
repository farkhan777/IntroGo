package main

import "fmt"

func main() {
	greeting := foo()
	fmt.Println(greeting)

	x := bar()
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	fmt.Println(x())
	fmt.Printf("%T\n", x())

	i := x()
	fmt.Println(i)
	fmt.Printf("%T\n", i)

	fmt.Println(bar()())
}

func foo() string {
	s := "Hello World!"
	return s
}

func bar() func() int {
	return func() int {
		return 451
	}
}
