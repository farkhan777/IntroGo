package main

import "fmt"

const (
	a = 45
	b = 45.78
	c = "Farkhan Hamzah Firdaus"
)

//const (
//	d int = 45
//	e float64 = 45.78
//	f string = "Farkhan Hamzah Firdaus"
//)

func main()  {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
