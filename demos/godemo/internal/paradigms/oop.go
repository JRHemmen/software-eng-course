package paradigms

import (
	"fmt"
	"math"
)

// Shape interface defines methods common to all geometric shapes.
// This allows for polymorphism where different shapes can be treated as the same type of object (Shape).
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle struct represents a circle with a specific radius.
// It 'implements' the Shape interface by providing the Area and Perimeter methods.
type Circle struct {
	Radius float64
}

// Area calculates the area of the Circle, demonstrating encapsulation of Circle's behavior.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter calculates the perimeter of the Circle.
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Rectangle struct represents a rectangle defined by its width and height.
// Like Circle, Rectangle also implements the Shape interface.
type Rectangle struct {
	Width, Height float64
}

// Area calculates the area of the Rectangle, another example of encapsulation.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter calculates the perimeter of the Rectangle.
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// ObjectOrientedProgramming demonstrates how OOP principles are applied in Go.
func ObjectOrientedProgramming() {

	// Instantiate Circle and Rectangle objects.
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 10, Height: 5}

	// Polymorphism in action: treating different shapes as the Shape interface.
	shapes := []Shape{circle, rectangle}
	for _, shape := range shapes {
		fmt.Printf("Area of %T: %.2f\n", shape, shape.Area())
		fmt.Printf("Perimeter of %T: %.2f\n\n", shape, shape.Perimeter())
	}

	// This loop demonstrates polymorphic behavior, where objects of different types
	// (Circle and Rectangle) are processed using a single interface type (Shape).
}
