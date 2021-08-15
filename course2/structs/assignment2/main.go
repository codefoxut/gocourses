package main

import "fmt"

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

type shape interface {
	getArea() float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println("area of the shape is ", s.getArea())
}

func main() {
	s1 := square{sideLength: 10.0}
	printArea(s1)

	t1 := triangle{base: 5.0, height: 10.0}
	printArea(t1)
}
