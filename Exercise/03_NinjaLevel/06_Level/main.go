package main

import (
	"fmt"
)

func main()  {
	var inputs string

	fmt.Print("Are you an asshole [yes/Y || no/N]? ")
	fmt.Scan(&inputs)
	if inputs == "yes" || inputs == "Y" {
		fmt.Println("You are an asshole")
	}
	fmt.Println("yeyy")
}
