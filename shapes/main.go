package main

import (
	"fmt"
	"math"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	side float64
}

func main() {
	t := triangle{
		base:   3.50,
		height: 9.35,
	}

	s := square{
		side: 6.76,
	}

	printArea(t)
	printArea(s)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	// Custom logic for generating an English greeting
	return t.base * t.base
}

func (s square) getArea() float64 {
	return math.Pow(s.side, 2)
}
