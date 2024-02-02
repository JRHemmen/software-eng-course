# Session 0: Programming Paradigms

## Functional Programming

- **Definition and Core Concepts**: Introducing functional programming as a paradigm centered around pure functions and avoiding shared state, mutable data, and side-effects.
- **Pure Functions**: Emphasizing functions that always produce the same output for the same input without any side effects.
- **Immutability**: Stressing the importance of immutable data structures that do not change once created.
- **First-Class and Higher-Order Functions**: Exploring the concepts of functions as first-class citizens and the use of higher-order functions that can take functions as parameters or return them as results.
- **Recursion**: Demonstrating how problems can be solved via functions that call themselves, often as a cleaner alternative to traditional looping mechanisms.

[See in Code](demos/godemo/internal/paradigms/functional.go)

## Object-Oriented Programming (OOP)

- **Introduction to OOP**: Presenting OOP as a paradigm that models real-world entities as objects with state (data) and behavior (methods).
- **Classes and Instances**: Clarifying the distinction between classes as blueprints for objects and instances as the actual objects created from these blueprints.
- **Attributes and Methods**: Delving into how data (attributes) and functions (methods) are encapsulated within objects.
- **Encapsulation**: Highlighting the encapsulation of data and methods within objects to hide the internal state and expose only necessary parts of the object interface.
- **Inheritance and Polymorphism**: Exploring inheritance as a mechanism for creating new classes from existing ones and polymorphism as the ability for objects of different classes to be treated as objects of a common superclass.

[See in Code](demos/godemo/internal/paradigms/oop.go)

## Compiled vs Interpreted Languages

- **Compiled Languages**: Languages like C and C++ are translated into machine code before execution, allowing for optimized performance. This catches errors early but requires recompilation for every change.
- **Interpreted Languages**: Languages like Python and JavaScript are translated on-the-fly during execution, offering flexibility and ease of use. This is a great choice for beginners to get started quickly but is ultimately slower than compiled languages.

[Compiled vs Interpreted Code Performance](https://medium.com/swlh/compiled-vs-interpreted-code-performance-e1a63299760b)

- Go code 2 orders of magnitude faster than Python code
- Pypy (JIT compiled Python) code 5x faster than native Python code

## Imperative vs Declarative Programming

- **Imperative Programming**: Focusing on the "how" of problem-solving, involving explicit instructions to achieve a desired state. Examples include procedural languages like C and assembly language, where control flow and state management are explicitly managed by the programmer.

- **Declarative Programming**: Concentrating on the "what" of problem-solving, abstracting the execution process. Examples include SQL for database queries and HTML for web page structure, where the desired outcome is specified without dictating the step-by-step process to achieve it.

## Memory Management

- **Data Types**: Exploring the different types of data (e.g., integers, floating-point numbers, strings, booleans) and their representation in memory.
- **Variables and Memory**: Understanding how data is stored and accessed in memory.
- **Stack and Heap**: Distinguishing between stack (for static memory allocation) and heap (for dynamic memory allocation) in the context of memory management.
- **Garbage Collection**: Exploring automated memory management techniques that reclaim memory allocated to objects no longer in use.
- **Memory Leaks**: Identifying and preventing situations where memory is not correctly released back to the system.
- **Pointers and References**: Delving into the use of pointers in low-level languages like C and references in higher-level languages to manipulate memory addresses directly.

[See in Code](demos/godemo/internal/paradigms/memory.go)

## Low-Level Programming

- **Assembly Language**: Introducing assembly language as a low-level programming language that is a direct representation of machine code.
- **Instruction Set Architecture (ISA)**: Exploring the instruction set architecture of a processor, which defines the set of instructions that a processor can execute. Examples include x86_64, ARM, and RISC-V.
- **Compilers, Assemblers, and Linkers**: Understanding the role of compilers, assemblers, and linkers in translating high-level code into executable binaries.