package main

import "testing"

func runSample(t *testing.T, k int64, l int64, r int64, z int64, x int, y int64, expect bool) {
	res := solve(k, l, r, z, x, y)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 8, 1, 10, 2, 6, 4, false)
}

func TestSample2(t *testing.T) {
	runSample(t, 8, 1, 10, 2, 6, 5, true)
}

func TestSample3(t *testing.T) {
	runSample(t, 9, 1, 10, 9, 2, 9, false)
}

func TestSample4(t *testing.T) {
	runSample(t, 20, 15, 25, 3, 5, 7, true)
}
