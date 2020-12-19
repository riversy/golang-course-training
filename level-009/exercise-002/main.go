package main

import "fmt"

type person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p *person) speak() {
	fmt.Printf(
		"I am %v %v of the %v years old.\n",
		p.FirstName,
		p.LastName,
		p.Age,
	)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {

	h := person{
		FirstName: "Lesly",
		LastName:  "Strike",
		Age:       28,
	}

	saySomething(&h)
}
