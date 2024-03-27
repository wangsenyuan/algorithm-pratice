package main

import "testing"

func runSample(t *testing.T, x int, y int, p int, q int, expect int) {
	res := solve(x, y, p, q)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x, y, p, q := 3, 10, 1, 2
	expect := 4
	runSample(t, x, y, p, q, expect)
}

func TestSample2(t *testing.T) {
	x, y, p, q := 7, 14, 3, 8
	expect := 10
	runSample(t, x, y, p, q, expect)
}
