package main

import "testing"

func runSample(t *testing.T, A, W []int, expect int64) {
	res := solve(len(A), len(W), A, W)

	if res != expect {
		t.Errorf("Sample %v %v, expect %d, but got %d", A, W, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 13, 7, 17}
	W := []int{1, 3}
	var expect int64 = 48
	runSample(t, A, W, expect)
}

func TestSample2(t *testing.T) {
	A := []int{10, 10, 10, 10, 11, 11}
	W := []int{3, 3}
	var expect int64 = 42
	runSample(t, A, W, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1000000000, 1000000000, 1000000000, 1000000000}
	W := []int{1, 1, 1, 1}
	var expect int64 = 8000000000
	runSample(t, A, W, expect)
}
