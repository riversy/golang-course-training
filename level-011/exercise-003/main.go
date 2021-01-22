package main

import "fmt"

type customerError struct {
	FirstName string
	LastLast  string
}

func (p customerError) Error() string {
	return fmt.Sprintf("error with %s %s", p.FirstName, p.LastLast)
}

func main() {
	c := customerError{"John", "Doe"}
	err := foo(c)
	if err != nil {
		panic(err)
	}
}

func foo(customer customerError) error {
	return customer
}
