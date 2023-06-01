package main

import "testing"

func runSample(t *testing.T, A []int, k int, expect int64) {
	res := solve(A, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 1, 1}
	k := 1
	var expect int64 = 3
	runSample(t, A, k, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 3}
	k := 10
	var expect int64 = 0
	runSample(t, A, k, expect)
}

func TestSample3(t *testing.T) {
	A := []int{5, 6}
	k := 2
	var expect int64 = 5
	runSample(t, A, k, expect)
}

func TestSample4(t *testing.T) {
	A := []int{6, 9, 4, 2, 1}
	k := 3
	var expect int64 = 6
	runSample(t, A, k, expect)
}
