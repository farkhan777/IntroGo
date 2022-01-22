package main

import "fmt"

func main() {
	foo()
	bar("Farkhan")
	fmt.Println(woo())
	s1 := wee("Moneypenny")
	fmt.Println(s1)
	m1, sure := mouse("Farkhan", "Hamzah Firdaus")
	fmt.Println(m1)
	fmt.Println(sure)
}

func foo() {
	fmt.Println("Hello from foo")
}

func bar(s string) {
	fmt.Println("Hello, ", s)
}

func woo() string {
	return "Hello from woo. "
}

func wee(st string) string {
	return fmt.Sprint("Hello from wee. ", st)
}

func mouse(firstName string, lastName string) (string, bool) {
	a := fmt.Sprint(firstName, " ", lastName, `, says "Hello"`)
	b := true
	return a, b
}
