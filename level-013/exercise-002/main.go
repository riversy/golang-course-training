package main

import (
	"fmt"
	"github.com/riversy/golang-course-training/level-013/exercise-002/quote"
	"github.com/riversy/golang-course-training/level-013/exercise-002/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
