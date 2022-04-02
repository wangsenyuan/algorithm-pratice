package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, A []int, B []int, expect float64) {
	res := solve(n, A, B)

	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	A := []int{-9, -21, -4, -43}
	B := []int{0, 0, 0, 0}
	var expect = 0.0
	runSample(t, n, A, B, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	A := []int{10, 14, 5, 9, 1}
	B := []int{4, 1, 3, 0, 0}
	var expect = 11.5
	runSample(t, n, A, B, expect)
}

func TestSample3(t *testing.T) {
	n := 7
	A := []int{-1, 101, 5, 63, -7, -88, 59}
	B := []int{0, 1, 6, 2, 3, 4, 5}
	var expect = 54.333333
	runSample(t, n, A, B, expect)
}
