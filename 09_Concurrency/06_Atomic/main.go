package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("Numb of CPUs:", runtime.NumCPU())
	fmt.Println("Numb of Goroutines:", runtime.NumGoroutine())

	var counter int64
	const gs = 100
	var wg sync.WaitGroup

	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("Numb of Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("Numb of Goroutines:", runtime.NumGoroutine())
	fmt.Println("Count:", counter)
}
