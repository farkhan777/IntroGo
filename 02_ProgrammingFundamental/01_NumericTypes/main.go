package main

import (
	"fmt"
	"runtime"
)

var x int8 = -128
var y float32 = 1234.51
var c float64 = 99999999999999999999999999999999999999999999999999999999999999999999999999957123812371982379183719387193871983719387198379128371982371982379182371983719837198237918237918379182739127391273912837918273981739146456456443434147162846218761284671824618246182467182467182416824124112412412412412412.1241241241
var z uint64 = 1

func main()  {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(c)
	fmt.Println(z)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}