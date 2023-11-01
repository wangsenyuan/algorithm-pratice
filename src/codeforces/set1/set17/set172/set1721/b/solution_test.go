package main

import "testing"

func runSample(t *testing.T, n, m, x, y, d int, expect int) {
	res := solve(n, m, x, y, d)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, x, y, d := 2, 4, 1, 3, 1
	expect := -1
	runSample(t, n, m, x, y, d, expect)
}
