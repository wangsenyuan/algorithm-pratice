package main

import "testing"

func runSample(t *testing.T, A []int, expect int64) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{10, 20, 30}
	var expect int64 = 3
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 1, 2}
	var expect int64 = 4
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 2, 3, 4, 5}
	var expect int64 = 10
	runSample(t, A, expect)
}
