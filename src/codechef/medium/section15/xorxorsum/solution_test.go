package main

import "testing"

func runSample(t *testing.T, A []int64, expect int64) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int64{1, 2, 3, 6}
	var expect int64 = 2
	runSample(t, A, expect)
}
