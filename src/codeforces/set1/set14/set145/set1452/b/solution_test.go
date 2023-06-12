package main

import "testing"

func runSample(t *testing.T, A []int, expect int64) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 2, 2}
	var expect int64 = 1
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 2, 3, 2}
	var expect int64
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{0, 3, 0}
	var expect int64 = 3
	runSample(t, A, expect)
}
