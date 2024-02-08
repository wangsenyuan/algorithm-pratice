package main

import "testing"

func runSample(t *testing.T, n int, a int, b int, expect int) {
	res := solve(n, a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, x, y := 7, 2, 3
	expect := 12
	runSample(t, n, x, y, expect)
}

func TestSample2(t *testing.T) {
	n, x, y := 12, 6, 3
	expect := -3
	runSample(t, n, x, y, expect)
}
