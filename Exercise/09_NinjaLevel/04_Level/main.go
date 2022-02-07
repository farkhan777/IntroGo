package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("Number of Goroutines:", runtime.NumGoroutine())
	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	var mt sync.Mutex

	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			mt.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mt.Unlock()
			wg.Done()
		}()
		fmt.Println("Number of Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Number of Goroutines:", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
}
