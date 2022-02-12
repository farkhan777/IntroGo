package main

import "fmt"

func main() {
	c := make(chan int)

	//go foo(c)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	//bar()
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("About to exit")
}

//func foo(c chan<- int)  {
//	for i := 0; i < 100; i++ {
//		c <- i
//	}
//	close(c)
//}
//
//func bar(c <-chan int)  {
//	for v := range c {
//		fmt.Println(v)
//	}
//}
