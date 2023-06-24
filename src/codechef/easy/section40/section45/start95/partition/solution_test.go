package main

import "testing"

func runSample(t *testing.T, A []int, expect int64) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, -2, 2}
	var expect int64 = 1
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{5, -2, 5}
	var expect int64 = -4
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{-1, -1, -1, -1}
	var expect int64 = 8
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{-7, -1, 5, 9, -1, -2, 7, 8}
	var expect int64 = 2
	runSample(t, A, expect)
}
