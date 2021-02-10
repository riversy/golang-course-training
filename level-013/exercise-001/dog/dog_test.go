package dog

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

type TestCase struct {
	HumanYears int
	DogYears   int
}

// getTestCases retrieves the list of test cases
func getTestCases() []TestCase {
	return []TestCase{
		{1, 7},
		{0, 0},
		{10, 70},
	}
}

//BenchmarkYears helps to understand how fast the function Years is.
func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(rand.Int())
	}
}

//BenchmarkYearsTwo helps to understand how fast the function YearsTwo is.
func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(rand.Int())
	}
}

//ExampleYears provides an example of the Years function usage.
func ExampleYears() {
	fmt.Println(Years(7))
	//Output: 49
}

//ExampleYearsTwo provides an example of the YearsTwo function usage.
func ExampleYearsTwo() {
	fmt.Println(YearsTwo(7))
	//Output: 49
}

//TestYears tests the Years function within the range of tests.
func TestYears(t *testing.T) {
	for _, testCase := range getTestCases() {
		dogYears := Years(testCase.HumanYears)
		assert.Equal(t, testCase.DogYears, dogYears, "Should be equal")
	}
}

//TestYearsTwo tests the YearsTwo function within the range of tests.
func TestYearsTwo(t *testing.T) {
	for _, testCase := range getTestCases() {
		dogYears := YearsTwo(testCase.HumanYears)
		assert.Equal(t, testCase.DogYears, dogYears, "Should be equal")
	}
}
