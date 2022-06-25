package main

import "testing"

func runSample(t *testing.T, n int, B int, x int, y int, expect int64) {
	res := solve(n, B, x, y)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, B, x, y := 5, 100, 1, 30
	var expect int64 = 15
	runSample(t, n, B, x, y, expect)
}

func TestSample2(t *testing.T) {
	n, B, x, y := 7, 1000000000, 1000000000, 1000000000
	var expect int64 = 4000000000
	runSample(t, n, B, x, y, expect)
}
