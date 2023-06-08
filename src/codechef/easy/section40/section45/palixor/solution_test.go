package main

import "testing"

func runSample(t *testing.T, A []int, expect int64) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{13, 27, 12, 26}
	var expect int64 = 8
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 2, 2}
	var expect int64 = 6
	runSample(t, A, expect)
}
