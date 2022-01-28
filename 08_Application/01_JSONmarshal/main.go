package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "Farkhan",
		Last:  "Hamzah Firdaus",
		Age:   21,
	}

	p2 := person{
		First: "Baby",
		Last:  "Sastra Diredja",
		Age:   25,
	}

	p3 := person{
		First: "Yusuf",
		Last:  "Bachtiar",
		Age:   21,
	}

	people := []person{p1, p2, p3}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(bs)
	fmt.Println(string(bs))
}
