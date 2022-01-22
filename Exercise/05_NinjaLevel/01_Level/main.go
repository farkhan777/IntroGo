package main

import "fmt"

type person struct {
	first        string
	last         string
	favIceFlavor []string
}

func main() {
	p1 := person{
		first:        "Farkhan",
		last:         "Hamzah Firdaus",
		favIceFlavor: []string{"Strawberry", "Vanilla", "Chocolate"},
	}

	p2 := person{
		first:        "Baby",
		last:         "Sastra Diredja",
		favIceFlavor: []string{"Vanilla"},
	}

	fmt.Println(p1.first, p1.last)
	for i, v := range p1.favIceFlavor {
		fmt.Println(i+1, v)
	}

	fmt.Println(p2.first, p2.last)
	for i, v := range p2.favIceFlavor {
		fmt.Println(i+1, v)
	}

}
