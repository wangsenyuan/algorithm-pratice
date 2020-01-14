package main

import "testing"

func runSample(t *testing.T, N, M, S int, H []int, expect int) {
	res := solve(N, M, S, H)

	if res != expect {
		t.Errorf("Sample %d %d %d %v, expect %d, but got %d", N, M, S, H, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, M, S := 5, 4, 10
	H := []int{10, 24, 30, 19, 40}
	expect := 2
	runSample(t, N, M, S, H, expect)
}

func TestSample2(t *testing.T) {
	N, M, S := 5, 4, 16
	H := []int{7, 16, 35, 10, 15}
	expect := 4
	runSample(t, N, M, S, H, expect)
}
