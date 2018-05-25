package math

import "testing"

func TestAverage(t *testing.T) {
	var v float64
	v = Average([]float64{1, 2, 3, 4, 5, 0})
	if v != 2.5 {
		t.Error("Expected 1.5, got ", v)
	}
}

func TestMax(t *testing.T) {
	var v float64
	v = Max([]float64{1, 2, 3, 4, 5, 0})
	if v != 5 {
		t.Error("Expected 5, got ", v)
	}
}

func TestMin(t *testing.T) {
	var v float64
	v = Min([]float64{1, 2, 3, 4, 5, 0})
	if v != 0 {
		t.Error("Expected 1.5, got ", v)
	}
}
