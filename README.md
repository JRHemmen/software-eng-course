# Software Engineering Course

## Programming Paradigms (Session 0)

### Compiled vs Interpreted Languages

- **Compiled Languages**: Detailing how languages like C and C++ are translated into machine code before execution, allowing for optimized performance.
- **Interpreted Languages**: Discussing languages like Python and JavaScript, which are translated on-the-fly during execution, offering flexibility and ease of use.

### Imperative vs Declarative Programming

- **Imperative Programming**: Focusing on the "how" of problem-solving, involving explicit instructions to achieve a desired state. Examples include procedural languages like C and assembly language, where control flow and state management are explicitly managed by the programmer.

- **Declarative Programming**: Concentrating on the "what" of problem-solving, abstracting the execution process. Examples include SQL for database queries and HTML for web page structure, where the desired outcome is specified without dictating the step-by-step process to achieve it.

### Memory Management

- **Data Types**: Exploring the different types of data (e.g., integers, floating-point numbers, strings, booleans) and their representation in memory.
- **Variables and Memory**: Understanding how data is stored and accessed in memory.
- **Stack and Heap**: Distinguishing between stack (for static memory allocation) and heap (for dynamic memory allocation) in the context of memory management.
- **Garbage Collection**: Exploring automated memory management techniques that reclaim memory allocated to objects no longer in use.
- **Memory Leaks**: Identifying and preventing situations where memory is not correctly released back to the system.
- **Pointers and References**: Delving into the use of pointers in low-level languages like C and references in higher-level languages to manipulate memory addresses directly.

### Functional Programming

- **Definition and Core Concepts**: Introducing functional programming as a paradigm centered around pure functions and avoiding shared state, mutable data, and side-effects.
- **Immutability**: Stressing the importance of immutable data structures that do not change once created.
- **First-Class and Higher-Order Functions**: Exploring the concepts of functions as first-class citizens and the use of higher-order functions that can take functions as parameters or return them as results.
- **Pure Functions**: Emphasizing functions that always produce the same output for the same input without any side effects.
- **Recursion**: Demonstrating how problems can be solved via functions that call themselves, often as a cleaner alternative to traditional looping mechanisms.

### Object-Oriented Programming (OOP)

- **Introduction to OOP**: Presenting OOP as a paradigm that models real-world entities as objects with state (data) and behavior (methods).
- **Classes and Instances**: Clarifying the distinction between classes as blueprints for objects and instances as the actual objects created from these blueprints.
- **Attributes and Methods**: Delving into how data (attributes) and functions (methods) are encapsulated within objects.
- **Encapsulation**: Highlighting the encapsulation of data and methods within objects to hide the internal state and expose only necessary parts of the object interface.
- **Inheritance and Polymorphism**: Exploring inheritance as a mechanism for creating new classes from existing ones and polymorphism as the ability for objects of different classes to be treated as objects of a common superclass.

## Web Development Basics

### Front-End Development

#### **Session 1: Introduction to Front-End Development**

- Overview of front-end web development and its role in creating user interfaces.
- Basic technologies: HTML, CSS, and JavaScript.
- Introduction to responsive design principles.

#### **Session 2: Deep Dive into HTML and CSS**

- Understanding the structure and semantics of HTML.
- Styling with CSS: selectors, box model, flexbox, and grid layouts.

#### **Session 3: JavaScript for Dynamic Web Pages**

- Basics of JavaScript programming and its role in the DOM.
- Event handling, DOM manipulation, and AJAX requests.
- Introduction to front-end frameworks and libraries (e.g., React, Vue.js).

### Back-End Development

#### **Session 4: Introduction to Back-End Development**

- Overview of back-end development and its role in web applications.
- Understanding server-client architecture and HTTP protocol.
- Introduction to server-side languages (e.g., Node.js, Python, Ruby).

#### **Session 5: Databases and Data Modeling**

##### **SQL Databases**

- **Relational Databases**: Understanding the relational model and its implementation in databases like MySQL and PostgreSQL.
- **ACID Transactions**: Understanding the ACID properties of transactions and their importance in ensuring data integrity.
- **CRUD Operations**: Introducing the CRUD operations (create, read, update, delete) for interacting with data in a database.
- **SQL Queries**: Exploring the SQL query language for retrieving and manipulating data in a relational database.
- **Database Design**: Understanding the process of designing a database schema and the use of entity-relationship diagrams (ERDs) for visualizing the relationships between entities.

##### **NoSQL**

- **Document Stores**: Introduction to document-oriented databases (e.g., MongoDB) and their query languages.
- **Key-Value Stores**: Understanding the use of key-value stores (e.g., Redis) for caching and fast data retrieval.
- **Time-Series Databases**: Introduction to time-series databases (e.g., InfluxDB) and their applications in monitoring and IoT.

#### **Session 6: API Design**

- RESTful API concepts and best practices.
- Introduction to GraphQL as an alternative to REST.
- Authentication and authorization techniques (e.g., OAuth, JWT).

## Further Topics

### Low-Level Programming

- **Assembly Language**: Introducing assembly language as a low-level programming language that is a direct representation of machine code.
- **Instruction Set Architecture (ISA)**: Exploring the instruction set architecture of a processor, which defines the set of instructions that a processor can execute. Examples include x86_64, ARM, and RISC-V.
- **Compilers, Assemblers, and Linkers**: Understanding the role of compilers, assemblers, and linkers in translating high-level code into executable binaries.

### Software Ecosystems and Licensing

- Comparison of proprietary vs open source software ecosystems
- Overview of major players: Microsoft (.NET, C#), Oracle (Java), GNU, FSF
- Understanding open source licenses and their implications
