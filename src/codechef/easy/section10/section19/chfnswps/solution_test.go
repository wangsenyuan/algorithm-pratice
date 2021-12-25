package main

import "testing"

func runSample(t *testing.T, n int, A []int, B []int, expect int64) {
	res := solve(n, A, B)

	if res != expect {
		t.Errorf("Sample %d %v %v, expect %d, but got %d", n, A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 1
	A := []int{1}
	B := []int{2}

	runSample(t, n, A, B, -1)
}

func TestSample2(t *testing.T) {
	n := 2
	A := []int{1, 2}
	B := []int{2, 1}

	runSample(t, n, A, B, 0)
}

func TestSample3(t *testing.T) {
	n := 2
	A := []int{1, 1}
	B := []int{2, 2}

	runSample(t, n, A, B, 1)
}
