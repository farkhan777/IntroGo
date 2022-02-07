package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Begin, Numb of CPUs:", runtime.NumCPU())
	fmt.Println("Begin, Numb of Goroutines:", runtime.NumGoroutine())
	wg.Add(2)
	go foo()
	go bar()
	zor()

	fmt.Println("Mid, Numb of CPUs:", runtime.NumCPU())
	fmt.Println("Mid, Numb of Goroutines:", runtime.NumGoroutine())
	wg.Wait()

	fmt.Println("About to exit")
	fmt.Println("End, Numb of CPUs:", runtime.NumCPU())
	fmt.Println("End, Numb of Goroutines:", runtime.NumGoroutine())
}

func foo() {
	fmt.Println("foo")
	wg.Done()
}

func bar() {
	fmt.Println("bar")
	wg.Done()
}

func zor() {
	fmt.Println("zor")
}
