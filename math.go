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

// Min function takes a slice of floats, sorts them, and returns the lowest value
func Min(xs []int) int {
	sort.Ints(xs)
	lowest := xs[0]
	return int(lowest)
}

// Max function takes a slice of floats, sorts them, and returns the highest value
func Max(xs []int) int {
	sort.Ints(xs)
	highest := xs[(len(xs) - 1)]
	return int(highest)
}
