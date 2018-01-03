package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, gates [][]int, expect float64) {
	res := solve(n, gates)
	if math.Abs(res-expect) > 0.00001 {
		t.Errorf("sample: %d %v, should give %f, but got %f", n, gates, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 1
	gates := [][]int{
		{0},
	}
	expect := 0.5
	runSample(t, n, gates, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	gates := [][]int{
		{0},
		{0},
		{1, 1, 2},
	}
	expect := 0.29289
	runSample(t, n, gates, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	gates := [][]int{
		{0},
		{0},
		{2, 1, 2},
	}
	expect := 0.70711
	runSample(t, n, gates, expect)
}

func TestSample4(t *testing.T) {
	n := 5
	gates := [][]int{
		{0},
		{0},
		{0},
		{2, 1, 2},
		{1, 3, 4},
	}
	expect := 0.40303
	runSample(t, n, gates, expect)
}
