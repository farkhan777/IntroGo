package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Numb of CPUs:", runtime.NumCPU())
	fmt.Println("Numb of Goroutines:", runtime.NumGoroutine())

	counter := 0
	const gs = 100
	var wg sync.WaitGroup

	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			//time.Sleep(time.Nanosecond)
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Numb of Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("Numb of Goroutines:", runtime.NumGoroutine())
	fmt.Println("Count:", counter)
}
