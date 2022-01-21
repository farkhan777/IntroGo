package main

import (
	"fmt"
)

func main()  {
	students := map[interface{}]interface{} {
		"Farkhan": 21,
		"Taufik": 20,
		"Alifia": 21,
		"Ajeng": 20,
		"Fahmi": 22,
	}
	fmt.Println(students)

	students["Nizar"] = 20
	fmt.Println(students)

	for v, k := range students {
		fmt.Println(v, k)
	}

	fmt.Println()

	a := map[interface{}]interface{} {
		"appId": 2,
		"fcmServerKey": "keyTestTest",
		"name": "com.app",
		"version": []int{1, 2, 3},
		"xyz": 3,
	}

	for i := range a {
		fmt.Println(i, a[i])
	}

	fmt.Println()

	for i, k := range a {
		fmt.Println(i, k)
	}
}
