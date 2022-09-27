package main

import "testing"

func runSample(t *testing.T, A []int64, B []int64, expect int64) {
	res := solve(A, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int64{2, 2, 5, 2, 2, 7}
	B := []int64{2, 5}
	expect := int64(7)
	runSample(t, A, B, expect)
}
