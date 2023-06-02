package main

import "testing"

func runSample(t *testing.T, A []int, K int, expect int64) {
	res := solve(A, K)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 1, 1}
	k := 2
	var expect int64 = 2
	runSample(t, A, k, expect)
}

func TestSample2(t *testing.T) {
	A := []int{3, 1, 2, 4}
	k := 2
	var expect int64 = 0
	runSample(t, A, k, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 1, 2, 2, 3}
	k := 2
	var expect int64 = 0
	runSample(t, A, k, expect)
}
