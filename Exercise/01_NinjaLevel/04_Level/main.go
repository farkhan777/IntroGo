package main

import "fmt"

type age int
var myAge age

func main()  {
	fmt.Println(myAge)
	fmt.Printf("%T\n", myAge)
	myAge = 21
	fmt.Println(myAge)
}
