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
		Saying: []string{"Hi, How are you?", "I am good"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))
}

func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		// return []byte{}, fmt.Errorf("There was an error in toJSON: %v", err)
		return []byte{}, fmt.Errorf("There was an error in toJSON: %v", err)
	}
	return bs, nil
}
