package main

import "fmt"

func main()  {
	var x [5]int
	fmt.Println(x)
	x[3] = 45
	fmt.Println(x)

	var y [58]string
	for i := 65; i <= 122; i++ {
		y[i-65] = string(i)
	}

	fmt.Println(y)
	fmt.Println(y[0])

}
