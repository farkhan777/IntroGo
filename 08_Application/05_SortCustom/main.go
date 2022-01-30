package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}

type ByAge []person

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

type ByName []person

func (bn ByName) Len() int           { return len(bn) }
func (bn ByName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn ByName) Less(i, j int) bool { return bn[i].Name < bn[j].Name }

func main() {
	p1 := person{"Farkhan", 21}
	p2 := person{"Taufik", 20}
	p3 := person{Name: "Alifia", Age: 21}
	p4 := person{Name: "Ajeng", Age: 19}

	people := []person{p1, p2, p3, p4}

	fmt.Println(people)

	fmt.Println("Sort by age")
	sort.Sort(ByAge(people))
	fmt.Println(people)

	fmt.Println()
	fmt.Println("Sort by name")
	sort.Sort(ByName(people))
	fmt.Println(people)
}
