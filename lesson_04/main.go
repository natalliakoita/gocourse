package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

// Method to print information about radius
func (c Circle) String() string {
	return fmt.Sprintf("Cirlce: radius %0.2f", c.radius)
}

// Method calculate area for struct Circle
func (c Circle) Area() float64 {
	r := math.Pow(c.radius, 2)
	return r * math.Pi
}

// Method calculate perimeter for struct Circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type Rectangle struct {
	height float64
	width  float64
}

// Method to print information about height and width
func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle with height %0.2f and width %0.2f", r.height, r.width)
}

// Method calculate area for struct Rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Method calculate perimeter for struct Rectangle
func (r Rectangle) Perimeter() float64 {
	a := r.width + r.height
	return 2 * a
}

type Shape interface {
	String() string
	Area() float64
	Perimeter() float64
}

func DescribeShape(s Shape) {
	fmt.Println(s)
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func main() {
	c := Circle{radius: 6}
	r := Rectangle{
		height: 8,
		width:  2,
	}
	DescribeShape(c)
	DescribeShape(r)
}
