package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, h int, S []int, expect float64) {
	res := solve(n, h, S)

	if math.Abs(res-expect) > 1e-7 {
		t.Errorf("Sample expect %.7f but got %.7f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, h := 3, 1
	S := []int{2, 1}
	expect := 1.0
	runSample(t, n, h, S, expect)
}

func TestSample2(t *testing.T) {
	n, h := 3, 1
	S := []int{1, 1}
	expect := -1.0
	runSample(t, n, h, S, expect)
}

func TestSample3(t *testing.T) {
	n, h := 3, 1
	S := []int{2, 2}
	expect := 0.6666667
	runSample(t, n, h, S, expect)
}

func TestSample4(t *testing.T) {
	n, h := 6, 3
	S := []int{5, 2, 3, 10, 5}
	expect := 0.3804348
	runSample(t, n, h, S, expect)
}
