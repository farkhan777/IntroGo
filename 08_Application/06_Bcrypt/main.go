package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
)

func main() {
	password := "farkhan12345678"
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(password)
	fmt.Println(bs)
	fmt.Println(string(bs))

	loginPassword := "farkhan12345678"
	// 	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPassword))
	err = bcrypt.CompareHashAndPassword([]byte(bs), []byte(loginPassword))
	if err != nil {
		fmt.Println("You can't login!")
		return
	}

	fmt.Println("You're logged in")
}
