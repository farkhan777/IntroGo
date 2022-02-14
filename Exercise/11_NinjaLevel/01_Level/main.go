package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First  string
	Last   string
	Saying []string
}

func main() {
	p1 := person{
		First:  "Farkhan",
		Last:   "Hamzah Firdaus",
		Saying: []string{"Hi, My name is Farkhan", "I'm 21 years old"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(bs))
}
