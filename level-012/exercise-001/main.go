package main

import (
	"fmt"
	"github.com/riversy/golang-course-training/level-012/exercise-001/dog"
)

func main() {
	years := 36
	dogYears, _ := dog.Years(years)
	fmt.Printf(
		"The dog would be a %v years old if it would be a %v years old human.\n",
		dogYears,
		years,
	)
}
