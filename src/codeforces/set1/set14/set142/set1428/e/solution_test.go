package main

import "testing"

func runSample(t *testing.T, k int, A []int, expect int64) {
	res := solve(k, A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 6
	A := []int{5, 3, 1}
	var expect int64 = 15
	runSample(t, k, A, expect)
}

func TestSample2(t *testing.T) {
	k := 4
	A := []int{19}
	var expect int64 = 91
	runSample(t, k, A, expect)
}

func TestSample3(t *testing.T) {
	k := 3
	A := []int{1000000}
	var expect int64 = 333333333334
	runSample(t, k, A, expect)
}
