package main

import "testing"

func runSample(t *testing.T, x int, y int, n int, expect int) {
	res := solve(x, y, n)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, x, y := 8, 1, 1
	expect := 4
	runSample(t, x, y, n, expect)
}

func TestSample2(t *testing.T) {
	n, x, y := 8, 1, 10
	expect := 8
	runSample(t, x, y, n, expect)
}
