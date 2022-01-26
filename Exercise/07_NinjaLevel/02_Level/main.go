package main

import "fmt"

type person struct {
	name string
}

func changeMe(p *person) {
	p.name = "Farkhan Hamzah Firdaus"
	//(*p).name = "Farkhan HF"
}

func main() {
	p1 := person{
		name: "Baby Sastra Diredja",
	}

	fmt.Println(p1)
	fmt.Println(p1.name)

	changeMe(&p1)
	fmt.Println(p1)
	fmt.Println(p1.name)
}
