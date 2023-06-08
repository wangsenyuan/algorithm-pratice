package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, A []int, exp []string, expect float64) {
	res := solve(A, exp)

	diff := math.Abs(res - expect)

	if diff/math.Max(1.0, expect) > 1e-7 {
		t.Fatalf("Sample expect %.8f, but got %.8f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{4, 3, 2, 1}
	exp := []string{
		"1 0.3",
		"3 1",
		"4 0.6",
	}
	expect := 0.6
	runSample(t, A, exp, expect)
}

func TestSample2(t *testing.T) {
	A := []int{4, 2, 1, 3, 5}
	exp := []string{
		"3 0.8",
		"4 0.6",
		"5 0.3",
	}
	expect := 0.72
	runSample(t, A, exp, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 3, 2, 4, 5, 6}
	exp := []string{
		"4 0.9",
		"5 0.3",
		"2 0.4",
		"6 0.7",
		"3 0.5",
	}
	expect := 0.989500
	runSample(t, A, exp, expect)
}

func TestSample4(t *testing.T) {
	A := []int{1, 2, 3, 4}
	exp := []string{
		"2 0.5",
		"4 0.1",
	}
	expect := 1.0
	runSample(t, A, exp, expect)
}
