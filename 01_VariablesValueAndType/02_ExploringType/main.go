package main

import "fmt"

var y = 42

// DECLARE the VARIABLE with the IDENTIFIER "z"
// is of TYPE string
// and I ASSIGN the VALUE "shaken, not stirred"

var z = "Shaken, not stirred"

var a = `Farkhan said, "not stirred"`

// this is a STATIC programming language
// a VARIABLE is DECLARED to hold a VALUE of a certain TYPE
// not a DYNAMIC programming language

func main()  {
	fmt.Println(y)
	fmt.Printf("Data type : %T\n\n", y)

	fmt.Println(z)
	fmt.Printf("Data type : %T\n\n", z)

	fmt.Println(a)
	fmt.Printf("Data type : %T\n\n", a)

	//z = 43
	//fmt.Println(z)
	//fmt.Printf("%T\n", z)

	b := "I want to eat burger"
	fmt.Println(b)
	fmt.Printf("%T\n\n", b)

	//b = 45
	//fmt.Println(b)
	//fmt.Printf("%T\n\n", b)
}
