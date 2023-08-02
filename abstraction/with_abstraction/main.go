package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circles struct {
	Radius float64
}

type Rectangles struct {
	Width  float64
	Height float64
}

func (c Circles) Area() float64 {
	return c.Radius * c.Radius * 3.14
}

func (r Rectangles) Area() float64 {
	return r.Height * r.Width
}

func main() {
	circle := Circles{Radius: 5}
	rectangle := Rectangles{Width: 10, Height: 10}

	shapes := []Shape{circle, rectangle}

	for _, shape := range shapes {
		fmt.Printf("Area %.2f\n", shape.Area())
	}
}
