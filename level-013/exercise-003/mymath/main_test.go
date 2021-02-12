package mymath

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestCenteredAvg(t *testing.T) {
	testCases := []struct {
		in  []int
		out float64
	}{
		{
			in:  []int{1, 4, 6, 8, 100},
			out: 6,
		},
		{
			in:  []int{0, 8, 10, 1000},
			out: 9,
		},
		{
			in:  []int{9000, 4, 10, 8, 6, 12},
			out: 9,
		},
		{
			in:  []int{123, 744, 140, 200},
			out: 170,
		},
		{
			in:  []int{22, 1},
			out: 0,
		},
		{
			in:  []int{22},
			out: 0,
		},
	}

	for _, testCase := range testCases {
		values := testCase.in
		assert.Equal(t, testCase.out, CenteredAvg(values), "Should be equals.")
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{4, 1, 6, 8, 100}))
	//Output: 6
}

func BenchmarkCenteredAvg(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		minLength := 3
		maxLength := 15
		length := int(math.Max(
			math.Min(
				float64(rand.Int()),
				float64(minLength),
			),
			float64(maxLength),
		))
		CenteredAvg(rand.Perm(length))
	}
}
