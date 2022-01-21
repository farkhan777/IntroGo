package main

import "fmt"

func main() {
	name := "farkhan"
	switch name {
	case "bambang", "farkhan", "aleex":
		fmt.Println("I am farkhan")
	case "faisal":
		fmt.Println("I am faisal")
	case "edo":
		fmt.Println("I am edo")
		fallthrough
	case "samuel":
		fmt.Println("I am samuel")
		fallthrough
	case "yanto":
		fmt.Println("I am yanto")
		fallthrough
	default:
		fmt.Println("I don't know")
	}

	fmt.Println()

	Idk()
}

func Idk()  {
	switch {
	case false:
		fmt.Println("I am")
	case 3 == 3:
		fmt.Println("I am 3")
	case 4 == 4:
		fmt.Println("I a")
		fallthrough
	case (7 == 9):
		fmt.Println("I am farkhan")
		fallthrough
	case (11 == 3):
		fmt.Println("I am yanto")
		fallthrough
	default:
		fmt.Println("I don't know")
	}
}