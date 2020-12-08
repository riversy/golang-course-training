package main

import "fmt"

func main() {
	favSport := "Foot ball"

	switch favSport {
	case "Football":
		fmt.Println("My favourite sport it Football.")
	case "Baseball":
		fmt.Println("My favourite sport it Baseball.")
	default:
		fmt.Println("No one of these is my favourite sport.")
	}
}
