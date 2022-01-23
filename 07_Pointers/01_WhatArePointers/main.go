package main

import "fmt"

func main() {
	a := 45
	fmt.Println(a)
	fmt.Println(&a) // & gives you the address

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	// b := a
	b := &a
	fmt.Println(b)
	fmt.Println(*b) // * gives you the value stored at an address when you have the address
	fmt.Println(*&a)

	// b = 43
	*b = 43
	fmt.Println(a)

	fmt.Println()

	c := &b
	fmt.Println(c)
	fmt.Println(&c)
	fmt.Println(*c)
	fmt.Println(**c) // pointing a pointer ? what ?

	fmt.Println()

	d := &c
	fmt.Println(d)
	fmt.Println(&d)
	fmt.Println(*d)
	fmt.Println(**d)
	fmt.Println(***d)

	fmt.Println()

	z := &a
	fmt.Println(z)
	fmt.Println(&z)
	fmt.Println(*z)

	*z = 99
	fmt.Println(a, b, c, d, z)
	fmt.Println(a, b, *c, **d, z)
	fmt.Println(a, *b, **c, ***d, *z)

}
