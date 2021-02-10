//Package word provides a set of functions to estimate the corpus properties.
package word

import (
	"regexp"
	"strings"
)

//UseCount generates the map with the number of words entering in the provided corpus.
func UseCount(s string) map[string]int {
	xs := strings.Fields(
		CleanText(s),
	)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

//Count calculates the number of words in the provided corpus.
func Count(s string) int {
	xs := strings.Fields(
		CleanText(s),
	)
	return len(xs)
}

//CleanText removes all punctuation and makes text lowercase.
func CleanText(s string) string {
	wordsOnlyMatch := regexp.MustCompile(`[^\w\s\-]`)
	return strings.ToLower(
		wordsOnlyMatch.ReplaceAllLiteralString(s, ""),
	)
}
