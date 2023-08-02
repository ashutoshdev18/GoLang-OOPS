package main

import (
	"fmt"
)

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func CircleArea(c Circle) float64 {
	return 3.14 * c.Radius * c.Radius
}

func RectangleArea(r Rectangle) float64 {
	return r.Width * r.Height
}
func main() {
	circle := Circle{Radius: 2.15}
	fmt.Println(CircleArea(circle))
	rectangle := Rectangle{Width: 5, Height: 10}
	fmt.Println(RectangleArea(rectangle))
}
