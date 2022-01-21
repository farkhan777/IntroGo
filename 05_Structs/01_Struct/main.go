package main

import "fmt"

type person struct {
	firstName string
	lastName string
	age int
	hobby []string
}

func main()  {
	p1 := person{
		firstName: "Farkhan",
		lastName: "Hamzah Firdaus",
		age: 21,
		hobby: []string{"Swimming", "Coding", "Reading", "Running"},
	}

	p2 := person{
		firstName: "Baby",
		lastName: "Sastra Diredja",
		age: 25,
		hobby: []string{"Reading", "Eating", "Drawing"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.firstName)
	fmt.Println(p1.hobby)
	fmt.Println(p1.hobby[0])
}
