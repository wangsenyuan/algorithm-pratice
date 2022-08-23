package main

import "testing"

func runSample(t *testing.T, A []int, expect int64) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 4}
	var expect int64 = 0
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 1, 1}
	var expect int64 = 2
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{-4, 2, -4, 2, -4, 2}
	var expect int64 = 15
	runSample(t, A, expect)
}
