package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 6}

	shapes := []Shape{circle, rectangle}

	for _, shape := range shapes {
		fmt.Printf("Luas: %f\n", shape.Area())
	}
}
