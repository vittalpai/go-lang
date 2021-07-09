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

func (c Circle) Permeter() float32 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Height float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

func (r Rectangle) Permeter() float32 {
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
	Permeter() float32
}

func PrintPermeter(sp ShapeWithPermeter) {
	fmt.Println("Permeter = ", sp.Permeter())
}

func main() {
	c := Circle{10}
	//fmt.Println("Area = ", c.Area())
	PrintArea(c)
	PrintPermeter(c)

	r := Rectangle{10, 12}
	//fmt.Println("Area =", r.Area())
	PrintArea(r)
	PrintPermeter(r)
}
