package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func getArea(shape Shape) {

	fmt.Println(shape.Area())
}

func main() {

	r := Rectangle{Width: 21, Height: 22}
	c := Circle{Radius: 30}

	getArea(r)
	getArea(c)
}
