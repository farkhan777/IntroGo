package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("err happened", err)
		log.Println("err happened", err)
		//log.Fatalln(err)
		//panic(err)
	}
}

// Println formats using the default formats for its operands and write to standard output
