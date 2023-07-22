package main

import "testing"

func runSample(t *testing.T, A []int, expect int64) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{-1, -2, -3, -4, -5}
	var expect int64 = -120
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{-1, -2, -3, 1, 2, -1}
	var expect int64 = 12
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 0, 5, 4, -4, 3, 1, 1, 0, -3, 2}
	var expect int64 = 720
	runSample(t, A, expect)
}
