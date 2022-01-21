package main

import "fmt"

func main()  {
	//i := 45
	if i := 45; i == 5 {
		fmt.Println("Hi, I am Farkhan")
	}
	fmt.Println("Hello")
	//fmt.Println(i)

	score := 100
	if score > 80 {
		fmt.Println("You got A")
	} else if score >= 75 && score <= 80 {
		fmt.Println("You got B+")
	} else if score >= 70 && score < 75 {
		fmt.Println("You got B")
	} else if score >= 60 && score < 70 {
		fmt.Println("You got C+")
	} else if score >= 50 && score < 60 {
		fmt.Println("You got C")
	} else if score >= 40 && score < 50 {
		fmt.Println("You got D")
	} else {
		fmt.Println("You got E")
	}

	for i := 0; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Println(i)
		}
	}
}
