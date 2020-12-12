package main

import "fmt"

func main() {
	message := struct {
		from    string
		to      string
		message string
	}{
		from:    "Igor",
		to:      "The World",
		message: "Hello, World!",
	}

	fmt.Println(message)
}
