package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	cr := (<-chan int)(c)
	cs := (chan<- int)(c)

	go func() {
		cs <- 42
	}()

	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)
}
