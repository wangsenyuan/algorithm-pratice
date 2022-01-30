package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, expect []float64) {
	res := solve(n, E)

	for i := 0; i < len(res); i++ {
		diff := math.Abs((res[i] - expect[i])) / math.Max(1, expect[i])
		if diff > 1e-6 {
			t.Errorf("Sample expect %f, but got %f", expect[i], res[i])
		}
	}
}

func TestSample1(t *testing.T) {
	n := 4
	options := [][]int{
		{1, 2, 1},
		{2, 3, 2},
		{3, 3, 1},
		{3, 4, 0},
		{2, 4, 1},
	}
	expect := []float64{
		0.000000000000,
		1.000000000000,
		2.000000000000,
		1.000000000000,
	}
	runSample(t, n, options, expect)
}

func TestSample2(t *testing.T) {
	n := 7
	options := [][]int{
		{1, 3, 3},
		{1, 2, 4},
		{1, 6, 5},
		{2, 3, 2},
		{3, 4, 3},
		{4, 5, 5},
		{5, 2, 1},
		{5, 3, 1},
		{4, 6, 2},
		{7, 1, 239},
	}
	expect := []float64{
		0.000000000000,
		4.000000000000,
		3.000000000000,
		4.500000000000,
		7.250000000000,
		4.250000000000,
		-1,
	}
	runSample(t, n, options, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	options := [][]int{
		{2, 3, 0},
		{3, 2, 0},
	}
	expect := []float64{0, -1, -1}
	runSample(t, n, options, expect)
}
