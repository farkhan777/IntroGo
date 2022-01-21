package main

import "fmt"

func main()  {
	p1 := struct {
		firstName string
		lastName string
		age int
		hobby []string
	}{
		firstName: "Farkhan",
		lastName: "Hamzah Firdaus",
		age: 21,
		hobby: []string{"Swimming", "Coding", "Reading", "Running"},
	}

	fmt.Println(p1)

	var x interface{} = 45.5
	var y interface{} = 35

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Println(x)
	fmt.Println(y)

	fmt.Println()
	y = x

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Println(x)
	fmt.Println(y)
}
