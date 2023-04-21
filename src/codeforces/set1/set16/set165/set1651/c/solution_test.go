package main

import "testing"

func runSample(t *testing.T, A []int, B []int, expect int64) {
	res := solve(A, B)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1000000000, 1000000000, 1000000000, 1000000000}
	B := []int{1, 1, 1, 1}
	var expect int64 = 1999999998
	runSample(t, B, A, expect)
}
