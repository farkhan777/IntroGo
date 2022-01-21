package main

import "fmt"

const (
	year1 = 2022 + iota
	year2 = 2022 + iota
	year3 = 2022 + iota
	year4 = 2022 + iota
)

func main()  {
	fmt.Println(year1)
	fmt.Println(year2)
	fmt.Println(year3)
	fmt.Println(year4)
}
