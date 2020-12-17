package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {
	(*p).first = "Changed"
}

func main() {
	p1 := person{
		first: "Dan",
		last:  "Brown",
		age:   33,
	}

	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}
