package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo() {
	fmt.Println("Foo")
	wg.Done()
}

func bar() {
	fmt.Println("Bar")
	wg.Done()
}

func main() {
	wg.Add(2)

	go foo()
	go bar()

	fmt.Println("Go routines started. Wait...")
	wg.Wait()
}
