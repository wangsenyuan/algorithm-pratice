package main

import "testing"

func runSample(t *testing.T, A, B []int, expect int) {
	res := solve(A, B)
	if expect != res {
		t.Errorf("Sample %v %v, expect %d, but got %d", A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 3, 4}
	B := []int{-20, -1, 18}
	expect := 2
	runSample(t, A, B, expect)
}
