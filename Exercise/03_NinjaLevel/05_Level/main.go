package main

import "fmt"

func main()  {
	for i := 10; i <= 100; i++ {
		fmt.Printf("When %v is devided by 4, the reminder (aka modulus) is %v\n", i, i%4)
	}
}
