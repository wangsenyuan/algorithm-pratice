package main

import "testing"

func runSample(t *testing.T, a, b, k int, A, B []int, expect int64) {
	res := solve(a, b, k, A, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a, b, k := 3, 4, 4
	A := []int{1, 1, 2, 3}
	B := []int{2, 3, 2, 4}
	var expect int64 = 4
	runSample(t, a, b, k, A, B, expect)
}

func TestSample2(t *testing.T) {
	a, b, k := 2, 2, 4
	A := []int{1, 1, 2, 2}
	B := []int{1, 2, 1, 2}
	var expect int64 = 2
	runSample(t, a, b, k, A, B, expect)
}
