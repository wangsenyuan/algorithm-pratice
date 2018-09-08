package main

import "testing"

func runSample(t *testing.T, N int, A []int, expect int) {
	res := solve(N, A)
	if int(res) != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", N, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N := 6
	A := []int{1, 2, 3, 4, 5, 6}
	runSample(t, N, A, 5)
}

func TestSample2(t *testing.T) {
	N := 9
	A := []int{0, 2, 1, 2, 1, 3, 0, 1, 0}
	runSample(t, N, A, 199)
}

func TestSample3(t *testing.T) {
	N := 7
	A := []int{0, 5, -5, 5, -5, 4, -4}
	runSample(t, N, A, 55)
}
