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
}
