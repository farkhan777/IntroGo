package main

import "fmt"

func main()  {
	lastFirst := map[string][]string {
		"bond_james": []string{`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
		"np_dr": []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	if _, ok := lastFirst["np_dr"]; ok {
		delete(lastFirst, "np_dr")
	}

	for i, v := range lastFirst {
		fmt.Println("This is the record for", i)
		for j, k := range v {
			fmt.Println("\t", j, k)
		}
	}
}
