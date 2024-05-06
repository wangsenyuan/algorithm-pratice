package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, cmds [][]int, expect []float64) {
	res := solve(cmds)

	if len(res) != len(expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}

	for i := 0; i < len(res); i++ {
		if math.Abs(res[i]-expect[i]) > 1e-6 {
			t.Fatalf("Sample expect %v, but got %v", expect, res)
		}
	}
}

func TestSample1(t *testing.T) {
	cmds := [][]int{
		{1, 3},
		{2},
		{1, 4},
		{2},
		{1, 8},
		{2},
	}
	expect := []float64{0, 0.5, 3.0}
	runSample(t, cmds, expect)
}

func TestSample2(t *testing.T) {
	cmds := [][]int{
		{1, 1},
		{1, 4},
		{1, 5},
		{2},
	}
	expect := []float64{2}
	runSample(t, cmds, expect)
}
