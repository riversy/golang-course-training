package math

import "testing"

// testPair collects a struct of inputs and outputs
type testPair struct {
	values  []float64
	average float64
}

// tests slice with inputs and outputs
var tests = []testPair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

//TestAverage test Average function using tests
func TestAverage(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}
