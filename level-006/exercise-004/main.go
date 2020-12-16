package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("I'm %v %v, %v years old.\n", p.first, p.last, p.age)
}

func main() {
	p1 := person{
		first: "Dan",
		last:  "Brown",
		age:   33,
	}

	p1.speak()
}
