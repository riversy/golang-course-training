package main

import "fmt"

func main() {
	x := [5]int{34, 43, 42, 1, 5}

	for _, v := range x {
		fmt.Println(v)
	}

	fmt.Printf("%T\n", x)
}
