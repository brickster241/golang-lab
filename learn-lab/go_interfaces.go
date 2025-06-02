package learnlab

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rect) perim() float64 {
	return 2 * r.height + 2 * r.width
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) diameter() float64 {
	return 2 * c.radius
}

func measure(g geometry) {
	fmt.Println("G :", g)
	fmt.Printf("Type of G : %T\n", g)
	fmt.Println("Area of G :", g.area())
	fmt.Println("Perimeter of G :", g.perim())
	fmt.Println("============================")
}

/*
	structs rect and circle are implementing geometry interface
*/
func Go_Interfaces() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	measure(r)
	measure(c)
}