package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Height float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float32 {
	return (2 * r.Height) + (2 * r.Width)
}

//utility func
type ShapeWithArea interface {
	Area() float32
}

func PrintArea(sa ShapeWithArea) {
	fmt.Println("Area = ", sa.Area())
}

type ShapeWithPermeter interface {
	Perimeter() float32
}

func PrintPerimeter(sp ShapeWithPermeter) {
	fmt.Println("Permeter = ", sp.Perimeter())
}

//interface composition

type Dimension interface {
	ShapeWithArea
	ShapeWithPermeter
}

func printDimensions(d Dimension) {
	PrintArea(d)
	PrintPerimeter(d)
}

func main() {
	c := Circle{10}
	//fmt.Println("Area = ", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	printDimensions(c)

	r := Rectangle{10, 12}
	//fmt.Println("Area =", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	printDimensions(r)
}
