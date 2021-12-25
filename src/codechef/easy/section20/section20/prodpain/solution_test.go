package main

import "testing"

func runSample(t *testing.T, A []int, expect int64) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3, 4}
	var expect int64 = 4
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 42, 69, 228, 2021}
	var expect int64 = 1041808
	runSample(t, A, expect)
}
