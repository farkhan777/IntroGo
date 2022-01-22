package main

import (
	"fmt"
	"math"
)

type circle struct {
	l float64
}

type square struct {
	r float64
}

func (c circle) area() float64 {
	return math.Pi * (c.l * c.l)
}

func (s square) area() float64 {
	return s.r * s.r
}

type shape interface {
	area() float64
}

func info(sp shape) {
	fmt.Println(sp.area())
}

func main() {
	cr := circle{
		l: 12.345,
	}

	sq := square{
		r: 15,
	}

	info(cr)
	info(sq)
}
