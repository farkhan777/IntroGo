package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("..\\IntroGo\\11_ErrorHandling\\01_CheckingErrors\\04\\names.txt")
	if err != nil {
		panic(err)
		return
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
		return
	}

	fmt.Println(string(bs))
}
