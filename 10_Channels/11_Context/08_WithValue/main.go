package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "language", "Go")
	foo(ctx, "language")

	ctx = context.WithValue(ctx, "dog", "Gaston")
	foo(ctx, "dog")

	foo(ctx, "color")
}

func foo(ctx context.Context, s string) {
	v := ctx.Value(s)
	if v != nil {
		fmt.Println("found value:", v)
		return
	}
	fmt.Println("key not found")
}
