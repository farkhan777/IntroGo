package main

import "fmt"

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("a norgate math error occured: %v", ce.info)
}

func main() {
	c1 := customErr{
		info: "need more coffee",
	}
	foo(c1)
}

func foo(e error) {
	fmt.Println("foo ran -", e, "\n")
}
