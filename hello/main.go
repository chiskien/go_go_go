package main

import "fmt"

func main() {
	var greeting = greet()
	greeting2 := greet()
	fmt.Println(greeting)
	fmt.Println(greeting2)
}

func greet() string {
	//return a simple greeting message
	return "Hello World"
}
