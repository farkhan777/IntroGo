package main

import "fmt"

type person struct {
	firstName string
	lastName string
	age int
	hobby []string
}

type secretAgent struct {
	person
	ltk bool
}

func main()  {
	s1 := secretAgent{
		person: person{
			firstName: "Farkhan",
			lastName: "Hamzah Firdaus",
			age: 21,
			hobby: []string{"Swimming", "Coding", "Reading", "Running"},
		},
		ltk: true,
	}

	fmt.Println(s1)
	fmt.Println(s1.firstName, s1.person.lastName, s1.age, s1.hobby[0], s1.ltk)
}
