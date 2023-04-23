package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, X []int, Y []int, R []float64, C []int) {
	rad, cnt := solve(X, Y)

	for i := 0; i < len(X); i++ {
		if math.Abs(rad[i]-R[i]) > 1e-2 || cnt[i] != C[i] {
			t.Fatalf("Sample result %v %v not correct", rad, cnt)
		}
	}
}

func TestSample1(t *testing.T) {
	X := []int{0, 0, 3}
	Y := []int{0, 0, 4}
	R := []float64{0, 0, 5.0}
	C := []int{1, 1, 2}
	runSample(t, X, Y, R, C)
}

func TestSample2(t *testing.T) {
	X := []int{5, 7, 0, 3, 4}
	Y := []int{3, 8, 9, 1, 4}
	R := []float64{1.41, 5, 6.4, 2.83, 1.41}
	C := []int{2, 4, 4, 2, 1}
	runSample(t, X, Y, R, C)
}
