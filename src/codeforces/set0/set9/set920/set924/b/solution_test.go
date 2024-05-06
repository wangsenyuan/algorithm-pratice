package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, u int, E []int, expect float64) {
	res := solve(u, E)

	if math.Abs(res-expect) > 1e-9 {
		t.Fatalf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	u := 4
	E := []int{1, 3, 5, 7}
	expect := 0.5
	runSample(t, u, E, expect)
}

func TestSample2(t *testing.T) {
	u := 8
	E := []int{10, 13, 15, 16, 17, 19, 20, 22, 24, 25}
	expect := 0.875
	runSample(t, u, E, expect)
}
