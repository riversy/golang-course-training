package main

import "fmt"

const x int = 42
const y = 43
const z = 43.43

func main() {
	fmt.Printf("%v\t%v\t%v\n", x, y, z)
	fmt.Printf("%T\t%T\t%T\n", x, y, z)
}
