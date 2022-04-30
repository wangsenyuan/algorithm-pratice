package main

import "testing"

func runSample(t *testing.T, A []int64, expect int64) {
	res := solve(len(A), A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int64{1, 1, 1, 1}
	var expect int64
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int64{1, 1, 1}
	var expect int64 = 1
	runSample(t, A, expect)
}
