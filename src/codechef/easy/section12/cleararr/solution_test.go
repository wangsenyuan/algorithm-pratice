package main

import "testing"

func runSample(t *testing.T, A []int, k int, x int, expect int64) {
	res := solve(len(A), k, x, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k, x := 2, 7
	A := []int{9, 10, 11, 12, 13}
	var expect int64 = 23
	runSample(t, A, k, x, expect)
}

func TestSample2(t *testing.T) {
	k, x := 0, 7
	A := []int{9, 9, 9, 9, 9}
	var expect int64 = 45
	runSample(t, A, k, x, expect)
}

func TestSample3(t *testing.T) {
	k, x := 2, 7
	A := []int{9, 1, 2, 3, 10}
	var expect int64 = 13
	runSample(t, A, k, x, expect)
}
