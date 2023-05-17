package main

import "testing"

func runSample(t *testing.T, A, B []int64, expect int64) {
	res := solve(A, B)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int64{2, 3, 5, 1, 2}
	B := []int64{4, 3, 6, 2, 3}
	var expect int64 = 3
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int64{0, 0, 0, 0}
	B := []int64{1, 2, 2, 1}
	var expect int64 = 2
	runSample(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := []int64{1, 2, 3}
	B := []int64{1, 2, 2}
	var expect int64 = -1
	runSample(t, A, B, expect)
}
