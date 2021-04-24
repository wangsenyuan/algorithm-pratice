package main

import "testing"

func runSample(t *testing.T, n int, C []int, expect int64) {
	res := solve(n, C)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	C := []int{4, 3, 2, 1, 4}
	var expect int64 = 19
	runSample(t, n, C, expect)
}
