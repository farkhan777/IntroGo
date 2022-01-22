package main

import "fmt"

func main() {
	//result := sum()
	result := sum("Farkhan")
	fmt.Println("The total is", result)
}

// func sum(numbers ...int) int {
// variadic parameter (numbers ...int) will always in the last position
func sum(s string, numbers ...int) int {
	fmt.Println(numbers)
	fmt.Printf("%T\n", numbers)

	sum := 0
	for i, v := range numbers {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding", v, "to the total which is now", sum)
	}
	fmt.Println(s)
	return sum
}
