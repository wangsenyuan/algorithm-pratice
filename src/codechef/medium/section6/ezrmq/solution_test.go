package main

import "testing"

func runSample(t *testing.T, k int, A []int, expect int64) {
	res := solve(k, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 3
	A := []int{1, 2, 3, 4, 5}
	var expect int64 = 14
	runSample(t, k, A, expect)
}
