# Learning Go

## Chapter 1: Meet Go

### Outlines:

> - Go's key features for modern software development
> - Go's history and simplicity-driven philosophy
> - Go's built-in tools for testing, benchmarking, and debugging
> - How hands-on projects help master Go
> - Go's versatility for backend and cloud applications

### 1.1 What is Go?

> Go is a programming language that was originally designed to solve various problems in large-scale software development in real-world

- Go is designed for large-scale software development by removing the constraint of memory handling, making it simple to run parallel
  pieces of code
- Go was built for **concurrency** and **networked servers**
- Go was designed to improve productivity during a time when multicore networked machines and large codebases were becoming the norm.

### 1.2 Why you should learn Go?

> Go is a versatile language designed for maintainability and readability. Optimal for backend software and great integration with
> modern cloud technologies.
> Go has `goroutines`, which is a safer and less costly way of dealing with parallel computing than threads.
> Threads rely on the operating systems, which is limited to size and power of CPU, whereas `goroutines` happen at the application's
> level.

- Go is not OOP and has no inheritance. Go supports most of the features via compositions and implicit interfaces

### 1.3 Comparing Go vs other languages

|                   | C++                                       | Python                      | Java                             | Go                                                                      |
|-------------------|-------------------------------------------|-----------------------------|----------------------------------|-------------------------------------------------------------------------|
| Design Philosophy | High-level OOP, procedural, multiparadigm | High Level OOP              | High Level OOP                   | **Procedural and data-oriented programming, support most OOP features** |
| Error Management  | Exceptions                                | Exception                   | Exceptions                       | Errors are values                                                       |
| Types             | Statically typed                          | Dynamic Typed               | Statically Typed                 | Statically Typed                                                        |
| Compilation       | Compiled                                  | Compiled at runtim          | Compiled                         | Compiled                                                                |
| Concurrency       | Memory-level or OS-level threads          | OS-Level Threads            | Memory-level or OS-level threads | Goroutines and channels                                                 |
| Interfaces        | Explicit                                  | Implicit and Explicit       | Explicit                         | Implicit                                                                |
| Memory release    | Full control                              | GC                          | GC                               | GC                                                                      |
| Use cases         | High Performance, Low overhead            | Excellent for Data Analysis | Well-suited for WebApplications  | Adapted for web APIs and Cloud                                          |
| Testing Tools     | Eternal Frameworks                        | Built-in native tools       | External Framwork: JUnit         | Built-in native tools                                                   |`       

## Chapter 2: Hello, Earth!

### Outlines:

- Writing to the standard output
- Testing writing to the standard output2
- Writing table-driven tests
- Using a hash table to hold key-value pairs
- Using flags to read command-line parameters

### What is in a name?

```go
package main
```

- Every `Go` file begins with the name of its package
- `Packages` are Go's way of organizing code, similar to modules or libraries in other languages.

### Capital Question

> Why `Println()` starts with an uppercase letter?

- Any symbol starting with a capital letter is exposed to external users of the package
- Anything else isn't accessible from outside the package
- Applied to variables, constants, functions, and types

### Internal and External testing

- Preparation Phase: 
- Execution Phase:
- Decision Phase:
- Teardown Phase: 

## Chapter 3: A Bookworm's Digest: Playing with Loops and Maps

## Chapter 4: A Log Story: Creating A Library

## Chapter 5: Gordle: Word game in Terminal

## Chapter 6: Money Converter

## Chapter 7: Caching with Generics

## Chapter 8: Gordle As A Service

## Chapter 9: Concurrent Maze Solver

## Chapter 10: Habits Tracker Using gRPC

## Chapter 11: HTML Templating with A gRPC Client

## Chapter 12: Go for other architectures