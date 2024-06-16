package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func (r triangle) getArea() float64 {
	return 0.5 * r.base * r.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	square := square{sideLength: 10}
	triangle := triangle{height: 10, base: 20}

	printArea(square)
	printArea(triangle)
}
