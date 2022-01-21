package main

import "fmt"

func main()  {
	favSport := "swimming"

	switch favSport {
	case "football":
		fmt.Println("Foot Ball")
	case "running":
		fmt.Println("Running")
	case "swimming":
		fmt.Println("Swimming")
	case "jerking":
		fmt.Println("Jerking")
	default:
		fmt.Println("You are an asshole")
	}
}
