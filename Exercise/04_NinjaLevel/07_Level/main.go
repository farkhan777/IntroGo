package main

import "fmt"

func main()  {
	group1 := []string{"James", "Bond", "Shaken, not stirred"}
	group2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	students := [][]string{group1, group2}

	fmt.Println(students)

	for i, v := range students {
		fmt.Println(i, v)
	}

	for i := 0; i < len(students); i++ {
		for j := 0; j <= len(students); j++{
			fmt.Println(students[i][j])
		}
	}

	for i, v := range students{
		fmt.Println("record: ", i)
		for j, val := range v {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}

	fmt.Println()

	for i := 0; i < len(students); i++ {
		fmt.Println("record: ", i)
		for j := 0; j <= len(students); j++{
			fmt.Printf("\t index position: %v \t value: \t %v \n",j, students[i][j])
		}
	}
}
