package main

import "fmt"

func main()  {
	myName := "Farkhan Hamzah Firdaus"
	fmt.Println(myName)
	x := []byte(myName)
	fmt.Println(x)

	for i := 0; i < len(myName); i++ {
		fmt.Printf("#%#U", myName[i])
	}

	fmt.Println("")

	for i := 0; i < len(myName); i++ {
		fmt.Printf("at index position %d we have hex %#x\n", i, myName[i])
	}
}
