package main

import "testing"

func runSample(t *testing.T, n int, a, b, c int, expect int) {
	res := solve(n, a, b, c)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, a, b, c := 10, 5, 5, 5
	expect := 9
	runSample(t, n, a, b, c, expect)
}
