package arg

import (
	"errors"
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

// Method to print information about radius
func (c Circle) String() string {
	return fmt.Sprintf("Cirlce: radius %0.2f", c.Radius)
}

// Method calculate area for struct Circle
func (c Circle) Area() (float64, error) {
	if c.Radius <= 0 {
		return 0, errors.New("invalid value radius")
	}
	return math.Pow(c.Radius, 2) * math.Pi, nil
}

// Method calculate perimeter for struct Circle
func (c Circle) Perimeter() (float64, error) {
	if c.Radius <= 0 {
		return 0, errors.New("invalid value radius")
	}
	return 2 * math.Pi * c.Radius, nil
}

type Rectangle struct {
	Height float64
	Width  float64
}

// Method to print information about height and width
func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle with height %0.2f and width %0.2f", r.Height, r.Width)
}

// Method calculate area for struct Rectangle
func (r Rectangle) Area() (float64, error) {
	if r.Height <= 0 {
		return 0, errors.New("invalid value radius")
	}
	if r.Width <= 0 {
		return 0, errors.New("invalid value radius")
	}
	return r.Width * r.Height, nil
}

// Method calculate perimeter for struct Rectangle
func (r Rectangle) Perimeter() (float64, error) {
	if r.Height <= 0 {
		return 0, errors.New("invalid value radius")
	}
	if r.Width <= 0 {
		return 0, errors.New("invalid value radius")
	}
	a := r.Width + r.Height
	return 2 * a, nil
}

type Shape interface {
	String() string
	Area() (float64, error)
	Perimeter() (float64, error)
}

func DescribeShape(s Shape) error {
	a, err := s.Area()
	if err != nil {
		return err
	}
	p, err := s.Perimeter()
	if err != nil {
		return err
	}
	fmt.Println(s)
	fmt.Printf("Area: %.2f\n", a)
	fmt.Printf("Perimeter: %.2f\n", p)
	return nil
}
