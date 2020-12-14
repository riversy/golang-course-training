package main

import "fmt"

func foo() int {
	return 42
}

func bar() (int, string) {
	return 9, "Miles"
}

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}
