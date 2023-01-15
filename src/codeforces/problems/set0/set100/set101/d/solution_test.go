package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, expect float64) {
	res := solve(n, E)

	if math.Abs(res-expect) > 1e-7 {
		t.Errorf("Sample expect %.7f, but got %.7f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	E := [][]int{
		{1, 2, 1},
	}
	expect := 1.0
	runSample(t, n, E, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	E := [][]int{
		{1, 3, 2},
		{4, 2, 1},
		{3, 2, 3},
	}
	expect := 4.333333333333334
	runSample(t, n, E, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2, 1},
		{1, 3, 1},
		{1, 4, 1},
		{1, 5, 1},
	}
	expect := 4.0
	runSample(t, n, E, expect)
}

func TestSample4(t *testing.T) {
	n := 4
	E := [][]int{
		{1, 2, 1},
		{1, 3, 1},
		{3, 4, 1},
	}
	expect := 2.6666666666666665
	runSample(t, n, E, expect)
}
