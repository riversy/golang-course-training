package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		if i%4 == 0 {
			fmt.Printf("The number %d is dividable to 4\n", i)
		}
	}
}
