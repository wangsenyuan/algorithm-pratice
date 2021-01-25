package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	A := []int{1, 0, 0, 1, 0, 1}
	expect := 5
	runSample(t, n, A, expect)
}
