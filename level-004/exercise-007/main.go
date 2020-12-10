package main

import "fmt"

func main() {
	data := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."},
	}

	for i, row := range data {
		fmt.Println(">>>", "Row", i+1, "<<<")
		for _, value := range row {
			fmt.Printf("\t- %v\n", value)
		}
	}
}
