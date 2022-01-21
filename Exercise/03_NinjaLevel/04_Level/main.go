package main

import "fmt"

func main()  {
	years := 2000
	for {
		years++
		if years > 2022 {
			break
		}
		fmt.Println(years)
	}
}
