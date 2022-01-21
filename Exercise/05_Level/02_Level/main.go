package main

import "fmt"

type person struct {
	first string
	last string
	favIceFlavor []string
}

func main()  {
	p1 := person{
		first: "Farkhan",
		last: "Hamzah Firdaus",
		favIceFlavor: []string{"Strawberry", "Vanilla", "Chocolate"},
	}

	p2 := person{
		first: "Baby",
		last: "Sastra Diredja",
		favIceFlavor: []string{"Vanilla"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	lastName := map[interface{}]person {
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(lastName["Hamzah Firdaus"])

	for _, v := range lastName {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for j, val := range v.favIceFlavor {
			fmt.Println(j, val)
		}
		fmt.Println("-------")
	}

}
