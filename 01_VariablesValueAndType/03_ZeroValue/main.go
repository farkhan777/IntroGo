package main

import "fmt"

var name string
var z int

func main()  {
	// DECLARE a variable to be of a certain TYPE
	// and then ASSIGN a VALUE of that type to the variable

	var x interface{}
	fmt.Println(x)
	fmt.Printf("%T\n\n", x)

	x = 5
	fmt.Println(x)
	fmt.Printf("%T\n\n", x)

	x = "asdas"
	fmt.Println(x)
	fmt.Printf("%T\n\n", x)

	fmt.Println("printing the value of y", name, "ending")
	fmt.Printf("%T\n\n", name)

	name = "Farkhan Hamzah Firdaus"

	fmt.Println("printing the value of y", name, "ending")
	fmt.Printf("%T\n\n", name)

	fmt.Println(z)
	fmt.Printf("%T", z)
}
