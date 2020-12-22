package main

import "testing"

func runSample(t *testing.T, A []int, B []int, expect int64) {
	res := solve(A, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{10, 5, 2, 5, 4}
	B := []int{3, 1, 2}
	var expect int64 = 5
	runSample(t, A, B, expect)
}
