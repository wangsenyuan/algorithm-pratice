package main

import "testing"

func runSample(t *testing.T, n int, A, B []int, expect int) {
	res := solve(n, A, B)

	if res != expect {
		t.Errorf("Sample %d %v %v, expect %d, but got %d", n, A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{1, 2, 3, 4, 5}
	B := []int{2, 3, 4, 5, 1}
	expect := 5
	runSample(t, n, A, B, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	A := []int{1, 3, 2, 4}
	B := []int{4, 2, 3, 1}
	expect := 2
	runSample(t, n, A, B, expect)
}
