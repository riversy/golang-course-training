package main

import "fmt"

func main() {
	aCountUp, aCountDown := getCounters()
	bCountUp, bCountDown := getCounters()

	fmt.Println("A(up):", aCountUp())
	fmt.Println("A(up):", aCountUp())
	fmt.Println("A(up):", aCountUp())
	fmt.Println("A(up):", aCountUp())
	fmt.Println("A(down):", aCountDown())

	fmt.Println("B(up):", bCountUp())
	fmt.Println("B(up):", bCountUp())
	fmt.Println("B(down):", bCountDown())

	fmt.Println("A(down):", aCountDown())
	fmt.Println("B(up):", bCountUp())
}

func getCounters() (func() int, func() int) {
	var count int

	return func() int {
			count++
			return count
		},
		func() int {
			count--
			return count
		}
}
