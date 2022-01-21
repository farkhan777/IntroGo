package main

import "fmt"

var x = 42

func main() {
	fmt.Println(x)
	fmt.Printf("%d\n", x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%x\n", x)
	fmt.Printf("%#x\n", x)
	x = 911
	fmt.Printf("%#x\n", x)
	fmt.Printf("%#x\t%b\t%x\n", x, x, x)

	y := fmt.Sprintf("%#x\t%b\t%x\n", x, x, x)
	fmt.Println(y)

	//z := 10
	c, _ := fmt.Printf("%#x\t%b\t%x\n", x, x, x)
	fmt.Println(c)
}
