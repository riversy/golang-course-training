package main

import "fmt"
import "github.com/riversy/golang-course-training/level-011/exercise-005/math"

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
}
