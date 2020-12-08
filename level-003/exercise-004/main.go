package main

import "fmt"

func main() {
	year := 1984
	age := 0
	for {
		if year > 2020 {
			break
		}
		fmt.Printf(
			"I was happy to live in the %v year in the age of %d years.\n",
			year,
			age,
		)
		year++
		age++
	}
}
