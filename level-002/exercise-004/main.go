package main

import "fmt"

var x int

func main() {
	x = 5
	fmt.Printf("%v\t%b\t%x\n", x, x, x)

	x = x << 1
	fmt.Printf("%v\t%b\t%x\n", x, x, x)
}
