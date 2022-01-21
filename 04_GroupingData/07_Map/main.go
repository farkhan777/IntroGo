package main

import "fmt"

func main()  {
	students := map[interface{}]int{
		"Farkhan": 21,
		"Taufik": 20,
		"Alifia": 21,
		"Ajeng": 20,
		"Fahmi": 22,
	}
	fmt.Println(students)
	fmt.Println(students["Farkhan"])
	fmt.Println(students["Rahmi"])

	v, ok := students["Farkhan"]
	fmt.Println(v, ok)

	if v, ok := students["Farkhan"]; ok {
		fmt.Println("This is the IF print", v)
	}

}
