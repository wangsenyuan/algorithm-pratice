package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, X []int, expect float64) {
	res, pos := solve(len(X), X)

	if math.Abs(res-expect) > 1e-6 {
		t.Fatalf("Sample expect %.7f, but got %.7f", expect, res)
	}
	n := len(X)
	var l int
	var eps float64 = 1e-6
	for i := 0; i < len(pos); i++ {
		a := search(0, n, func(j int) bool {
			return float64(X[j]) >= pos[i]-res
		})
		if a > l {
			t.Fatalf("Sample result %v, not correct, as it not cover %d", pos, l)
		}
		// a <= l
		b := search(0, n, func(j int) bool {
			return float64(X[j]) > pos[i]+res+eps
		})
		l = b
	}
	if l != n {
		t.Fatalf("Sample result %v, not correct, as not it only extends to %d", pos, l)
	}
}

func TestSample1(t *testing.T) {
	X := []int{1, 2, 3, 4}
	expect := 0.500000
	runSample(t, X, expect)
}

func TestSample2(t *testing.T) {
	X := []int{10003, 10004, 10001, 10002, 1}
	expect := 0.5
	runSample(t, X, expect)
}

func TestSample3(t *testing.T) {
	X := []int{26, 21, 20, 91, 22, 28, 92, 62, 47, 69}
	expect := 11.0
	runSample(t, X, expect)
}
