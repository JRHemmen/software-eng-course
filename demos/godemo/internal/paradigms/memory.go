package paradigms

import "fmt"

func Memory() {
	// Data Types
	fmt.Println("Data Types:")
	// Demonstrating various data types in Go.
	var integer int = 42
	var floatingPoint float64 = 3.1415
	var str string = "Hello, Go!"
	var boolean bool = true
	fmt.Printf("Integer: %d, Floating Point: %f, String: %s, Boolean: %t\n\n",
		integer, floatingPoint, str, boolean)
	fmt.Println("Each data type in Go, such as int, float64, string, and bool, occupies a certain amount of memory and has a specific use case.")

	// Variables and Memory
	fmt.Println("Variables and Memory:")
	// Demonstrating how variables are stored in memory.
	fmt.Printf("Integer address: %p, Floating Point address: %p\n\n", &integer, &floatingPoint)
	fmt.Println("Variables in Go are stored in memory, and each variable's memory address can be accessed using the '&' operator.")

	// Stack and Heap
	fmt.Println("Stack and Heap:")
	// Demonstrating stack vs. heap memory allocation.
	stackVar := "Stored on the stack" // Allocated on the stack.
	heapVar := new(string)            // Allocated on the heap.
	*heapVar = "Stored on the heap"
	fmt.Printf("Stack variable: %s, Heap variable: %s\n\n", stackVar, *heapVar)
	fmt.Println("Variables can be allocated on the stack (for short-lived variables) or on the heap (for variables that need to persist). The 'new' keyword allocates on the heap.")

	// Pointers and References
	fmt.Println("Pointers and References:")
	// Demonstrating the use of pointers.
	ptr := &integer // Pointer to an integer.
	fmt.Printf("Pointer to integer: %p, Dereferenced pointer: %d\n\n", ptr, *ptr)
	fmt.Println("Pointers hold the memory address of a variable. Dereferencing a pointer ('*ptr') accesses the variable's value at that address.")

	// Garbage Collection
	fmt.Println("Garbage Collection:")
	// Explaining automatic memory management in Go.
	fmt.Println("Go has a built-in garbage collector that automatically frees up unused memory, helping to prevent memory leaks and manage heap-allocated variables.")

	// Memory Leaks
	fmt.Println("Memory Leaks:")
	// Discussing potential memory leaks.
	fmt.Println("Memory leaks occur when allocated memory is not properly released back to the system. In Go, this can happen if pointers or references to heap-allocated objects are not let go.")
}
