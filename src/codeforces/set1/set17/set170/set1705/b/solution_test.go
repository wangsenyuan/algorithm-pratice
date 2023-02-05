package main

import "testing"

func runSample(t *testing.T, A []int, expect int64) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 0, 0}
	var expect int64 = 3
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{0, 2, 0, 2, 0}
	var expect int64 = 5
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{2, 0, 3, 0, 4, 6}
	var expect int64 = 11
	runSample(t, A, expect)
}
