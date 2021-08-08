package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, persons [][]int, expect float64) {
	res := solve(n, persons)

	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	persons := [][]int{
		{50, 50},
		{50, 60},
		{70, 100},
		{100, 60},
	}
	var expect float64 = 0.842105
	runSample(t, n, persons, expect)
}
