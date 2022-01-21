package main

import "fmt"

const (
	a = 45
)

func main()  {
	fmt.Println(a)
	fmt.Printf("%d\t%b\t%#x",a, a, a)
	b := a << 1
	fmt.Println()
	fmt.Printf("%d\t%b\t%#x",b, b, b)
}
