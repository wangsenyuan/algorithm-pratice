package main

import "testing"

func runSample(t *testing.T, A []int, C []int, expect int64) {
	res := solve(A, C)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 1, 4, 1, 5}
	C := []int{1, 4, 5, 5, 8, 99}
	var expect int64 = 20
	runSample(t, A, C, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1}
	C := []int{4, 1}
	var expect int64 = 4
	runSample(t, A, C, expect)
}
