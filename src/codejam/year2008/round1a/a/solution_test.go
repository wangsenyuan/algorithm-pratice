package main

import "testing"

func runSample(t *testing.T, A, B []int, expect int64) {
	res := solve(A, B)
	if res != expect {
		t.Errorf("Sample %v %v, expect %d, but got %d", A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 3, -5}
	B := []int{-2, 4, 1}
	runSample(t, A, B, -25)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 3, 4, 5}
	B := []int{1, 0, 1, 0, 1}
	runSample(t, A, B, 6)
}

func TestSample3(t *testing.T) {
	A := []int{1, -5, 3}
	B := []int{-2, 1, 4}
	runSample(t, A, B, -25)
}
