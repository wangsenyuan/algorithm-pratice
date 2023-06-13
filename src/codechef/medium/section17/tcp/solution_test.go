package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, tv []int, vv []int, x int, expect float64) {
	res := solve(x, tv, vv)

	if math.Abs(res-expect) > 1e-9 {
		t.Fatalf("Sample expect %.12f, but got %.12f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x := 5
	tv := []int{2, 3, 4, 5}
	vv := []int{11, 31, 69, 131}
	expect := 173.750000
	runSample(t, tv, vv, x, expect)
}

func TestSample2(t *testing.T) {
	x := 8
	tv := []int{12, 13, 14, 15}
	vv := []int{-480, -693, -960, -1287}
	expect := 80.0
	runSample(t, tv, vv, x, expect)
}
