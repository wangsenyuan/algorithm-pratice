package main

import "testing"

func runSample(t *testing.T, n int, A, B []int, expect int) {
	res := solve(n, A, B)

	if res != expect {
		t.Errorf("Sample %d %v, %v, expect %d, but got %d", n, A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 7
	A := []int{10, 10, 10, 10, 10, 3, 3}
	B := []int{1, 2, 1, 2, 1, 1, 1}
	expect := 4
	runSample(t, n, A, B, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	A := []int{1, 1, 1}
	B := []int{2, 3, 4}
	expect := 1
	runSample(t, n, A, B, expect)
}
