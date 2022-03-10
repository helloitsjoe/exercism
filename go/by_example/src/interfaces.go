package by_example

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
	perim() float64
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (r Rectangle) perim() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c Circle) perim() float64 {
	return 2 * c.radius * math.Pi
}

func measure(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func Interfaces() {
	r := Rectangle{width: 3, height: 4}
	c := Circle{radius: 4}

	measure(r)
	measure(c)
}
