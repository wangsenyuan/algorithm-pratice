package main

import "testing"

func runSample(t *testing.T, A []int, k int, expect int64) {
	res := solve(A, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{5, 0, -4}
	k := 2
	var expect int64 = 6
	runSample(t, A, k, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, -2, 0, 3, -4, 5}
	k := 4
	var expect int64 = 7
	runSample(t, A, k, expect)
}

func TestSample3(t *testing.T) {
	A := []int{0, 0, 0}
	k := 1000000000
	var expect int64 = 1000000001
	runSample(t, A, k, expect)
}

func TestSample4(t *testing.T) {
	A := []int{-7, -3, 8, 12, 0}
	k := 9
	var expect int64 = -1
	runSample(t, A, k, expect)
}
