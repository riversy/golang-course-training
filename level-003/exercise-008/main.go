package main

import "fmt"

func main() {
	for _, i := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {

		switch {
		case i%2 == 0:
			fmt.Printf("The number %d is even\n", i)
		case i%2 != 0:
			fmt.Printf("The number %d is odd\n", i)
		}
	}
}
