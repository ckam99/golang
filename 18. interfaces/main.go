package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Square struct {
	size float64
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	x, y, radius float64
}

func (s *Square) Area() float64 {
	return s.size * s.size
}

func (r *Rectangle) Area() float64 {
	return r.height * r.width
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func GetArea(s Shape) float64 {
	return s.Area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 10}
	rect := Rectangle{width: 23, height: 90}
	square := Square{size: 15}
	fmt.Println("Circle area:", GetArea(&circle))
	fmt.Println("Rectangle area:", GetArea(&rect))
	fmt.Println("Square area:", GetArea(&square))
}
