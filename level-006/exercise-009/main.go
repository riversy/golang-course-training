package main

import "fmt"

func getAFunc() func() {
	f := func() {
		fmt.Println("Hello! I'm the func from the func - A!")
	}
	return f
}

func getBFunc() func() {
	f := func() {
		fmt.Println("Hello! I'm the func from the func - B!")
	}
	return f
}

func runAnyFunc(f func()) {
	f()
}

func main() {
	fx := []func(){
		getAFunc(),
		getBFunc(),
	}

	for _, f := range fx {
		runAnyFunc(f)
	}
}
