package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Hi", p.first, p.last, "are you", p.age, "years old?")
}

func main() {
	p1 := person{
		first: "Farkhan",
		last:  "Hamzah Firdaus",
		age:   21,
	}

	p1.speak()
}
