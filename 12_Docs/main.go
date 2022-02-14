package main

import (
	"./dog"
	"fmt"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Farkhan",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
}
