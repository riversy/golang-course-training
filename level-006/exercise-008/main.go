package main

import "fmt"

func getAFunc() func() {
	f := func() {
		fmt.Println("Hello! I'm the func from the func!")
	}
	return f
}

func main() {
	f := getAFunc()
	f()
}
