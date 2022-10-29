package main

import "testing"

func runSample(t *testing.T, n int, W []int, S []int, expect int) {
	res := int(solve(n, W, S))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	W := []int{1, 2, 3, 4}
	S := []int{4, 3, 2, 1}
	expect := 7
	runSample(t, n, W, S, expect)
}
