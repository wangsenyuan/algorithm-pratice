package main

import "testing"

func runSample(t *testing.T, A []int, B []int, expect int64) {
	res := solve(A, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{10, 3, 7, 9, 1, 19, 2}
	B := []int{1, 4, 6}
	var expect int64 = 43

	runSample(t, A, B, expect)
}
