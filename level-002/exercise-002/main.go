package main

import "fmt"

func main() {
	x := 12
	y := 21
	a := x == y
	b := x <= y
	c := x >= y
	d := x != y
	e := x < y
	f := x > y

	fmt.Println("x :=", x)
	fmt.Println("y :=", y)
	fmt.Println("a := x == y", a)
	fmt.Println("b := x <= y", b)
	fmt.Println("c := x >= y", c)
	fmt.Println("d := x != y", d)
	fmt.Println("e := x < y", e)
	fmt.Println("f := x > y", f)
}
