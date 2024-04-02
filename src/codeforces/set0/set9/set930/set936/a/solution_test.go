package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, k int, d int, x int, expect float64) {
	res := solve(k, d, x)

	if math.Abs(res-expect)/math.Max(1, expect) >= 1e-7 {
		t.Fatalf("Sample expect %.8f, but got %.8f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k, d, x := 3, 2, 6
	expect := 6.5
	runSample(t, k, d, x, expect)
}

func TestSample2(t *testing.T) {
	k, d, x := 4, 2, 20
	expect := 20.0
	runSample(t, k, d, x, expect)
}

func TestSample3(t *testing.T) {
	k, d, x := 403, 957, 1000000000000000000
	expect := 1407352941176470446.0
	runSample(t, k, d, x, expect)
}
