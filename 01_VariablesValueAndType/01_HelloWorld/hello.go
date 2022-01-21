package main

import "fmt"

func main() {
	fmt.Println("Hello Go! ")
	foo()
	fmt.Println("Something more")

	for i := 0; i < 100; i++ {
		if i % 2 == 0 {
			fmt.Println(i)
		}
	}

	testVar()
}

func foo() {
	fmt.Println("This is foo!")
}

func testVar()  {
	n, _ := fmt.Println("Hello,",42," playground", 42, true)
	fmt.Printf("This is bite -> %d", n)
}
