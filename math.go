package math

import "sort"

// Average function takes a slice of floats and calulates the average
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

// Max function takes a slice of floats, sorts them, and returns the highest value
func Max(xs []float64) float64 {
	sort.Float64s(xs)
	lowest := xs[0]
	return float64(lowest)
}

// Min functoin takes a slice of floats, sorts them, and returns the lowest value
func Min(xs []float64) float64 {
	sort.Float64s(xs)
	highest := xs[(len(xs) - 1)]
	return float64(highest)
}
