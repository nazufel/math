package math

import "testing"

type testPairsAverage struct {
	values  []float64
	average float64
}

var testAverage = []testPairsAverage{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

func TestAverage(t *testing.T) {
	for _, pair := range testAverage {
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

type testPairsMax struct {
	values []int
	max    int
}

var testMax = []testPairsMax{
	{[]int{1, 2, 3, 4, 5}, 5},
	{[]int{1, 1, 1, 1, 1}, 1},
	{[]int{-1, 1}, 1},
}

func TestMax(t *testing.T) {
	for _, pair := range testMax {
		v := Max(pair.values)
		if v != pair.max {
			t.Error(
				"For", pair.values,
				"expected", pair.max,
				"got", v,
			)
		}
	}
}

type testPairsMin struct {
	values []int
	min    int
}

var testMin = []testPairsMin{
	{[]int{1, 2, 3, 4, 5}, 1},
	{[]int{1, 1, 1, 1, 1}, 1},
	{[]int{-1, 1}, -1},
}

func TestMin(t *testing.T) {
	for _, pair := range testMin {
		v := Min(pair.values)
		if v != pair.min {
			t.Error(
				"For", pair.values,
				"expected", pair.min,
				"got", v,
			)
		}
	}
}
