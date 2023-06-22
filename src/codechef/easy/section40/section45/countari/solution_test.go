package main

import "testing"

func runSample(t *testing.T, A []int, expect int64) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 5, 3, 6, 3, 4, 10, 4, 5, 2}
	var expect int64 = 9
	runSample(t, A, expect)
}
