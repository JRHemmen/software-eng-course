package paradigms

import "fmt"

func Memory() {
	// Data Types
	var integer int = 42
	var floatingPoint float64 = 3.1415
	var str string = "Hello, Go!"
	var boolean bool = true

	fmt.Println("Data Types:")
	fmt.Printf("Integer: %d, Floating Point: %f, String: %s, Boolean: %t\n\n", integer, floatingPoint, str, boolean)

	// Variables and Memory
	fmt.Println("Variables and Memory:")
	fmt.Printf("Integer address: %p, Floating Point address: %p\n\n", &integer, &floatingPoint)

	// Stack and Heap
	fmt.Println("Stack and Heap:")
	stackVar := "Stored on the stack"
	heapVar := new(string)
	*heapVar = "Stored on the heap"
	fmt.Printf("Stack variable: %s, Heap variable: %s\n\n", stackVar, *heapVar)

	// Pointers and References
	fmt.Println("Pointers and References:")
	ptr := &integer
	fmt.Printf("Pointer to integer: %p, Dereferenced pointer: %d\n\n", ptr, *ptr)

}
