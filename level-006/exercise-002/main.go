package main

import "fmt"

func foo(x ...int) int {
	var sum int
	for _, i := range x {
		sum += i
	}
	return sum
}

func bar(xi []int) int {
	return foo(xi...)
}

func main() {
	fmt.Println(
		bar(
			[]int{1, 2, 3, 4, 5},
		),
	)
}
