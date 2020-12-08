package main

import "fmt"

func main() {
	for code := 65; code <= 90; code++ {
		fmt.Println(code)
		for i := 1; i <= 3; i++ {
			fmt.Printf("\t%#U\n", code)
		}
	}
}
