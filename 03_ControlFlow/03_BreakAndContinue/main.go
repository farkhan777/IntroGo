package main

import "fmt"

func main()  {
	x := 1
	for {
		x++
		if x > 100 {
			break
		}

		if x%2 != 0 {
			continue
		}
		fmt.Println(x)
	}

	Ascii()
}

func Ascii()  {
	for i := 22; i <= 122; i++ {
		fmt.Printf("%v\t%#U\n", i, i)
	}
}
