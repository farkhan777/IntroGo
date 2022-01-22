package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type secretAgent struct {
	person
	ltk bool
}

// func (r receiver) identifier(parameters) (returns(s)) { code }
func (s secretAgent) speak() {
	fmt.Println("I am", s.firstName, s.lastName)
}

func main() {
	sa1 := secretAgent{
		person: person{
			firstName: "Farkhan",
			lastName:  "Hamzah Firdaus",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			firstName: "Baby",
			lastName:  "Sastra Diredja",
		},
		ltk: true,
	}

	fmt.Println(sa1)

	sa1.speak()
	sa2.speak()
}
