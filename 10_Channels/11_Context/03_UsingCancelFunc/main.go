package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Println("context:\t\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("----------")

	ctx, cancel := context.WithCancel(ctx)

	fmt.Println("context:\t\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Printf("cancel type:\t%T\n", cancel)
	fmt.Println("----------")

	cancel()

	fmt.Println("context:\t\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Printf("cancel type:\t%T\n", cancel)
	fmt.Println("cancel:\t\t\t", cancel)
	fmt.Println("----------")
}
