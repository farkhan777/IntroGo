package main

import "fmt"

func main()  {
	s := "F"
	fmt.Println(s)

	bs := []byte(s)
	fmt.Println(bs)

	b := bs[0]
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Printf("%b\n", b)
	fmt.Printf("%#X\n", b)
}
