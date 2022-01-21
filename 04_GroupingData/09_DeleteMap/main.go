package main

import "fmt"

func main()  {
	students := map[interface{}]interface{} {
		"Farkhan": 21,
		"Taufik": 20,
		"Alifia": 21,
		"Ajeng": 20,
		"Fahmi": 22,
	}
	fmt.Println(students)

	delete(students, "Fahmi")
	fmt.Println(students)

	if v, ok := students["Ajeng"]; ok {
		fmt.Println("value: ", v)
		delete(students, "Ajeng")
	}

	fmt.Println(students)

	for v, k := range students {
		fmt.Println(v, k)
	}
}
