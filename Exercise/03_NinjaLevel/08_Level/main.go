package main

import "fmt"

func main()  {
	option := 3

	switch {
	case option == 1 :
		fmt.Println("Opsi 1")
	case option == 2:
		fmt.Println("Opsi 2")
	case option == 3:
		fmt.Println("Opsi 3")
		//fallthrough
	default:
		fmt.Println("Tidak memilih")
	}
}
