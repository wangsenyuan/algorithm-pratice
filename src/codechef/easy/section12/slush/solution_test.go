package main

import "testing"

func runSample(t *testing.T, M int, C []int, N int, D []int, F []int, B []int, expect int) {
	res, _ := solve(M, C, N, D, F, B)
	if res != expect {
		t.Errorf("Sample %d %v %d %v %v %v, expect %d, but got %d", M, C, N, D, F, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, M := 5, 3
	C := []int{1, 2, 3}
	D := []int{2, 2, 2, 1, 1}
	F := []int{6, 10, 50, 10, 7}
	B := []int{3, 7, 3, 75, 4}

	runSample(t, M, C, N, D, F, B, 33)
}
