package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("..\\IntroGo\\11_ErrorHandling\\01_CheckingErrors\\04\\names.txt")
	if err != nil {
		panic(err)
		return
	}
	defer f.Close()

	r := strings.NewReader("Farkhan Hamzah Firdaus")

	io.Copy(f, r)
}
