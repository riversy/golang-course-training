package word

import (
	"fmt"
	"github.com/riversy/golang-course-training/level-013/exercise-002/quote"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type CountTestCase struct {
	in  string
	out int
}

type UseCountTestCase struct {
	in  string
	out map[string]int
}

type CleanTextTestCase struct {
	in  string
	out string
}

var countTestCases []CountTestCase
var useCountTestCases []UseCountTestCase
var cleanTextTestCases []CleanTextTestCase

func init() {
	countTestCases = []CountTestCase{
		{
			in:  "",
			out: 0,
		},
		{
			in:  "Lorem ipsum dolor sit amet.",
			out: 5,
		},
		{
			in:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam vel.",
			out: 10,
		},
	}

	useCountTestCases = []UseCountTestCase{
		{
			in:  "",
			out: map[string]int{},
		},
		{
			in: "Lorem ipsum dolor sit amet. Lorem ipsum.",
			out: map[string]int{
				"lorem": 2,
				"ipsum": 2,
				"dolor": 1,
				"sit":   1,
				"amet":  1,
			},
		},
	}

	cleanTextTestCases = []CleanTextTestCase{
		{
			in:  "",
			out: "",
		},
		{
			in:  "Lorem ipsum dolor sit amet. Lorem ipsum.",
			out: "lorem ipsum dolor sit amet lorem ipsum",
		},
		{
			in:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam vel.",
			out: "lorem ipsum dolor sit amet consectetur adipiscing elit etiam vel",
		},
		{
			in:  "-1",
			out: "-1",
		},
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkCleanText(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CleanText(quote.SunAlso)
	}
}

func ExampleUseCount() {
	fmt.Println(UseCount("Lorem ipsum dolor sit amet."))
	//Output: map[amet:1 dolor:1 ipsum:1 lorem:1 sit:1]
}

func ExampleCount() {
	fmt.Println(Count("Lorem ipsum dolor sit amet."))
	//Output: 5
}

func ExampleCleanText() {
	fmt.Println(CleanText("Lorem ipsum dolor sit amet."))
	//Output: lorem ipsum dolor sit amet
}

func TestCount(t *testing.T) {
	for _, testCase := range countTestCases {
		wordsCount := Count(testCase.in)
		assert.Equal(t, testCase.out, wordsCount, "Word count expected to be the same.")
	}
}

func TestUseCount(t *testing.T) {
	for _, testCase := range useCountTestCases {
		wordsMap := UseCount(testCase.in)
		assert.Equal(
			t,
			true,
			reflect.DeepEqual(wordsMap, testCase.out),
			fmt.Sprintf(
				"Word Maps expected to be the same. Expcected: %v. Actual: %v",
				testCase.out,
				wordsMap,
			),
		)
	}
}

func TestCleanText(t *testing.T) {
	for _, testCase := range cleanTextTestCases {
		cleanText := CleanText(testCase.in)
		assert.Equal(t, testCase.out, cleanText, "Strings expected to be the same.")
	}
}
