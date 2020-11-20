package main

import (
	"fmt"
)

type shape interface {
    getName() string
	getArea() float64
}

type triangle struct {
    name string
	height float64
	base   float64
}

type square struct {
    name string
	sideLength float64
}

func (t triangle) getName() string {
	return t.name
}

func (s square)  getName() string {
	return s.name
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println("Area:", s.getName(), s, s.getArea())
}

func main() {
	t := triangle{name:"Triangle", height: 10.0, base: 5.0}
	s := square{name:"Square",sideLength: 4}

	printArea(t)
	printArea(s)
}
