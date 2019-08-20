package main

import (
	"fmt"
	"math"
)

type react struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r react) area() float64 {
	return r.height * r.width
}

func (r react) perim() float64 {
	return 2*r.width + 2*r.height
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	r := react{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}

type geometry interface {
	area() float64
	perim() float64
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())

}
