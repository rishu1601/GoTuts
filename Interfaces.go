package main

import (
	"fmt"
	"math"
)

//To implement an interface,just give definitions to all functions defined in the interface
type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("Area:", g.area()) //Calls area based on the type of struct passed to it
	fmt.Println("Perimeter:", g.perimeter())
}
func main() {
	r := rect{width: 5, height: 4} //Create a rectangle
	c := circle{radius: 3}         //Create a circle

	measure(r) //Calling rectangle with interface geometry
	measure(c)
}
