package main

import "testing"

func runSample(t *testing.T, n int, a1 int, x int, y int, m int, k int, expect int) {
	res := solve(n, a1, x, y, m, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, a1, x, y, m, k := 5, 8, 2, 3, 100, 2
	expect := 1283
	runSample(t, n, a1, x, y, m, k, expect)
}
