package main

import (
	. "fmt"
	. "math"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * Pi
}

func main() {
	r1 := Rectangle{3, 6}
	c1 := Circle{5}

	Println("Area r1 is:", r1.area())
	Println("Area c1 is:", c1.area())
}