package paradigms

import (
	"fmt"
)

// AreaRectangle calculates the area of a rectangle given its width and height.
func AreaRectangle(width, height float64) float64 {
	return width * height
}

// PerimeterRectangle calculates the perimeter of a rectangle given its width and height.
func PerimeterRectangle(width, height float64) float64 {
	return 2 * (width + height)
}

// HigherOrderFunction takes a function 'f' as an argument along with two float64 operands 'x' and 'y'.
// The function 'f' is expected to have a signature that takes two float64 arguments and returns a float64 result.
// This demonstrates how higher-order functions can operate on functions with specific signatures.
func HigherOrderFunction(f func(float64, float64) float64, x, y float64) float64 {
	return f(x, y)
}

func FunctionalProgramming() {
	// Pure Functions
	fmt.Println("Pure Functions:")
	fmt.Println("A pure function's output is only determined by its input values, without observable side effects.")
	fmt.Printf("Area of rectangle with width 10 and height 5: %.2f\n", AreaRectangle(10, 5))

	// Function Signature
	fmt.Println("\nFunction Signature:")
	fmt.Println("A function signature defines the function's parameters and return type. For example, 'func(a, b float64) float64' is a signature for functions taking two float64 arguments and returning a float64.")

	// Immutability
	fmt.Println("\nImmutability:")
	const circleType string = "circle"
	fmt.Println("Immutable variables, once set, cannot be altered. This is a key principle in functional programming to avoid side effects.")
	fmt.Printf("A circle is always a %s, and this cannot be changed.\n", circleType)

	// Higher Order Functions
	fmt.Println("\nHigher Order Functions:")
	add := func(a, b float64) float64 { return a + b }
	result := HigherOrderFunction(add, 10, 5)
	fmt.Printf("Using a higher-order function to add 10 and 5 gives us: %f\n", result)
	fmt.Println("Higher-order functions can take any function matching the required signature (two float64 arguments, one float64 return).")

	// Anonymous Functions
	fmt.Println("\nAnonymous/Lambda Functions:")
	func(message string) {
		fmt.Printf("Anonymous functions don't have a name and can access variables from the enclosing function: %s\n", message)
	}("Go is awesome!")

	// Recursion
	fmt.Println("\nRecursion:")
	fmt.Println("Recursion is a technique where a function calls itself in order to solve smaller instances of the same problem.")
	fmt.Printf("Factorial of 5: %d\n", Factorial(5))
}

func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}
