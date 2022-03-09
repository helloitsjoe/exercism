package by_example

import "fmt"

type Rect struct {
	width, height int
}

func (r *Rect) area() int {
	return r.width * r.height
}

func (r Rect) perim() int {
	return 2 * (r.width + r.height)
}

func Methods() {
	r := Rect{width: 10, height: 5}

	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &Rect{width: 3, height: 2}

	fmt.Println("pointer area:", rp.area())
	fmt.Println("pointer perim:", rp.perim())
}
