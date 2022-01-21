package main

import "fmt"

func main()  {
	k := 0
	for i := 65; i < 91; i++ {
		k++
		fmt.Printf("%d\n", k)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\t%#U\n", i)
		}
	}
}
