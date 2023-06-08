package main

import "testing"

func runSample(t *testing.T, A []int, expect int64) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 4, 3}
	var expect int64 = 2
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 1, 1, 1}
	var expect int64 = 4
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{5, 6, 1, 3, 2, 9, 8, 1, 2, 4}
	var expect int64 = 15
	runSample(t, A, expect)
}
