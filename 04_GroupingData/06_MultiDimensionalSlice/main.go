package main

import "fmt"

func main()  {
	fruits := []string{"apple", "mango", "orange", "grape", "watermelon"}
	fmt.Println(fruits)
	colors := []string{"red", "green", "orange", "purple", "darkgreen"}
	fmt.Println(colors)
	items := [][]string{fruits, colors}
	fmt.Println(items)

}
