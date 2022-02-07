package main

import "fmt"

type person struct {
	name string
	age  int
}

type human interface {
	speak()
}

func (s *person) speak() {
	fmt.Println("Hello")
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		name: "Farkhan Hamzah Firdaus",
		age:  21,
	}

	p2 := person{
		name: "Baby Sastra Driedja",
		age:  25,
	}

	saySomething(&p1)
	saySomething(&p2)

	p1.speak()
}
