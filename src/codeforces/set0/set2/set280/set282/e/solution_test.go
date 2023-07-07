package main

import "testing"

func runSample(t *testing.T, A []int64, expect int64) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int64{1, 2}
	var expect int64 = 3
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int64{1, 2, 3}
	var expect int64 = 3
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int64{1000, 1000}
	var expect int64 = 1000
	runSample(t, A, expect)
}
