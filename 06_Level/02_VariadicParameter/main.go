package main

import "fmt"

func main()  {
	//numbs := []int{2, 3, 4, 5 ,6, 7, 8, 9}
	//result := sum(numbs)
	// or use
	//numbs := []int{2, 3, 4, 5 ,6, 7, 8, 9}
	//result := sum(numbs...)
	// or use
	result := sum(2, 3, 4, 5 ,6, 7, 8, 9)
	fmt.Println("The total is", result)
}
// func sum(numbers []int) int {}

// This is using variadic parameter (numbers ...int)
func sum(numbers ...int) int {
	fmt.Println(numbers)
	fmt.Printf("%T\n", numbers)

	sum := 0
	for i, v := range numbers {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding", v, "to the total which is now", sum)
	}
	return sum
}
