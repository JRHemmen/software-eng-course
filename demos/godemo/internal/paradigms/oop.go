package paradigms

import (
	"fmt"
	"math"
)

// Shape interface defines methods common to all shapes.
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle struct represents a circle with a radius.
type Circle struct {
	Radius float64
}

// Area method calculates the area of a Circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter method calculates the perimeter of a Circle.
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Rectangle struct represents a rectangle with width and height.
type Rectangle struct {
	Width, Height float64
}

// Area method calculates the area of a Rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter method calculates the perimeter of a Rectangle.
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// HigherOrderFunction takes a Shape interface and a function that operates on Shape.
func HigherOrderFunction(s Shape, f func(Shape) float64) float64 {
	return f(s)
}

func ObjectOrientedProgramming() {

	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 10, Height: 5}

	// Demonstrating polymorphism: Circle and Rectangle are both Shapes.
	shapes := []Shape{circle, rectangle}
	for _, shape := range shapes {
		fmt.Printf("Area of %T: %.2f\n", shape, shape.Area())
		fmt.Printf("Perimeter of %T: %.2f\n\n", shape, shape.Perimeter())
	}
}
