package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius int
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(float64(c.Radius), 2)
}

type Rectangle struct {
	Length int
	Width  int
}

func (r Rectangle) Area() float64 {
	return float64(r.Length * r.Length)
}

func PrintArea(shape Shape) {
	fmt.Printf("The area of shape is %f \n", shape.Area())
}

func main() {
	circle := Circle{2}
	PrintArea(circle)

	rectange := Rectangle{6, 8}
	PrintArea(rectange)
}
