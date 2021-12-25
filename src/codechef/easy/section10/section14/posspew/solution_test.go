package main

import "testing"

func runSample(t *testing.T, A []int, k int, expect int64) {
	res := solve(A, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{0, 1, 0, 1, 0, 0}
	k := 1
	var expect int64 = 6
	runSample(t, A, k, expect)
}

func TestSample2(t *testing.T) {
	A := []int{0, 0, 0, 2, 0, 0, 0, 5}
	k := 2
	var expect int64 = 23
	runSample(t, A, k, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 2, 3}
	k := 10
	var expect int64 = 66
	runSample(t, A, k, expect)
}
