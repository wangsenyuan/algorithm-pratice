package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, x []float64, T float64, expect float64) {
	res := solve(n, x, T)
	if math.Abs(res-expect) > 0.0001 {
		t.Errorf("sampel: %d %v %f, should give %f, but got %f", n, x, T, expect, res)
	}
}

func TestSample1(t *testing.T) {
	x := []float64{1, 2}
	n := 2
	T := float64(4)
	expect := 2.0
	runSample(t, n, x, T, expect)
}

func TestSample2(t *testing.T) {
	x := []float64{1, 2}
	n := 2
	T := float64(2)
	expect := 0.5
	runSample(t, n, x, T, expect)
}

func TestSample3(t *testing.T) {
	x := []float64{0, 0.5, 0.6, 2.75}
	n := 4
	T := float64(1)
	expect := 1.4
	runSample(t, n, x, T, expect)
}
