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

func (s secretAgent) speak() {
	fmt.Println("I am", s.firstName, s.lastName, " - the secretAgent speak")
}

func (p person) speak() {
	fmt.Println("I am", p.firstName, p.lastName, " - the person speak")
}

type human interface {
	speak()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into barrrrr", h.(person).firstName)
	case secretAgent:
		fmt.Println("I was passed into barrrrr", h.(secretAgent).firstName)
	}
	fmt.Println("I was passed into bar", h)
}

func main() {
	s1 := secretAgent{
		person: person{
			firstName: "Farkhan",
			lastName:  "Hamzah Firdaus",
		},
		ltk: true,
	}

	p1 := person{
		firstName: "Indah",
		lastName:  "Nala",
	}

	s1.speak()
	p1.speak()

	fmt.Println(p1)

	bar(s1)
	bar(p1)
}
