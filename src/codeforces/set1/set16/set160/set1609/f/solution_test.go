package main

import "testing"

func runSample(t *testing.T, A []int64, expect int64) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int64{1, 2, 3, 4, 5}
	var expect int64 = 9
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int64{0, 5, 7, 3, 9, 10, 1, 6, 13, 7}
	var expect int64 = 18
	runSample(t, A, expect)
}
