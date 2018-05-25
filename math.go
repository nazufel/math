package math

import "sort"

func Average(xs []float64) float64 {
  total := float64(0)
  for _, x := range xs {
    total += x
  }
  return total / float64(len(xs))
}

func Max(xs []float64) float64 {
  total := float64(0)
  for _, x := range xs {
    total += x
  }
  sort.Float64s(xs)
  lowest := xs[0]
  return total / float64(lowest)
}


func Min(xs []float64) float64 {
  total := float64(0)
  for _, x := range xs {
    total += x
  }
  sort.Float64s(xs)
  highest := xs[(len(xs)-1)]
  return total / float64(highest)
}
