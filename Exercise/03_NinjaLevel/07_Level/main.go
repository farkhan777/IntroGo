package main

import (
	"fmt"
	"strings"
)

func main()  {
	var inputs string

	fmt.Print("Are you an asshole [yes/y || no/n]? ")
	fmt.Scan(&inputs)
	inputs = strings.ToLower(inputs)
	if inputs == "yes" || inputs == "y" {
		fmt.Println("You are an asshole")
	} else if inputs == "no" || inputs == "n" {
		fmt.Println("You are not an asshole")
	}
	fmt.Println("yeyy")
}
