package main

import (
	"fmt"
	"math/rand"
)

var arr [10]int

func main()  {
	for i := 0; i < len(arr); i++ {
		numb := rand.Intn(100)
		if contains(arr, numb){
			continue
		} else {
			arr[i] = numb
		}
	}
	fmt.Printf("%d ", arr)

	fmt.Println()

	for _, v := range arr {
		fmt.Printf("%d ", v)
	}

	fmt.Printf("\n%T", arr)
}

func contains(s [10]int, str int) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
