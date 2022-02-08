package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area(num int) float64
	perim(num int) float64
}

type rect struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

func (r rect) area(num int) float64 {
	fmt.Println("Rect Area: ", num)
	return r.width * r.height
}

func (r rect) perim(num int) float64 {
	fmt.Println("Rect Perim: ", num)
	return 2 * (r.width + r.height)
}

func (c circle) area(num int) float64 {
	fmt.Println("Circle Radius: ", num)
	return math.Pi * c.radius * c.radius
}

func (c circle) perim(num int) float64 {
	fmt.Println("Circle Perim: ", num)
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area(1))
	fmt.Println(g.perim(2))
}

func main() {
	r := rect{width: 10, height: 5}
	c := circle{radius: 11}

	measure(r)
	measure(c)
}
