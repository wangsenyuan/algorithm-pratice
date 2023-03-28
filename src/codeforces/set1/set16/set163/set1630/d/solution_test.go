package main

import "testing"

func runSample(t *testing.T, A []int, B []int, expect int64) {
	res := solve(A, B)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{0, 6, -2, 1, -4, 5}
	B := []int{1, 2}
	var expect int64 = 18
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, -1, 1, -1, 1, -1, 1}
	B := []int{2}
	var expect int64 = 5
	runSample(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := []int{-1000000000, -1000000000, -1000000000, -1000000000, -1000000000}
	B := []int{1}
	var expect int64 = 5000000000
	runSample(t, A, B, expect)
}
