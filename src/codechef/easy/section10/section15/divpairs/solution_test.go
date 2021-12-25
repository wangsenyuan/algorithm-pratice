package main

import "testing"

func runSample(t *testing.T, A []int, expect int64) {
	res := solve(len(A), A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3}
	var expect int64 = 3
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 3, 5, 6}
	var expect int64 = 2
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{3, 9, 14, 6, 10}
	var expect int64 = 6
	runSample(t, A, expect)
}
