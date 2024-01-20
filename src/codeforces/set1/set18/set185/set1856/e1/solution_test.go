package main

import "testing"

func runSample(t *testing.T, n int, p []int, expect int) {
	res := solve(n, p)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	p := []int{1, 1, 3, 3}
	expect := 4
	runSample(t, n, p, expect)
}
