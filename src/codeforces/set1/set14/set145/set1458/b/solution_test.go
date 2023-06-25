package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, A []int, B []int, expect []float64) {
	res := solve(A, B)

	for i := 0; i < len(expect); i++ {
		if math.Abs(expect[i]-res[i]) > 1e-9 {
			t.Fatalf("Sample expect %v, but got %v", expect, res)
		}
	}
}

func TestSample1(t *testing.T) {
	A := []int{6, 6, 10}
	B := []int{5, 5, 2}
	expect := []float64{7, 11, 12}
	runSample(t, A, B, expect)
}
