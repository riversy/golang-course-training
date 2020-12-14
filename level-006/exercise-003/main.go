package main

import "fmt"

func foo() {
	defer fmt.Println("I'm a deferred call from Foo() :D")
	fmt.Println("Hello, World from Foo()!")
}

func main() {
	defer fmt.Println("I'm a deferred call from Main() :D")
	fmt.Println("Hello, World from Main()!")

	foo()
}
