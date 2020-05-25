package main

import "testing"

func runSample(t *testing.T, N, m, x, y int, A []int, expect int) {
	res := solve(N, m, x, y, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, m, x, y := 3, 1, 1, 5
	A := []int{1, 6, 9}
	expect := 38
	runSample(t, N, m, x, y, A, expect)
}

func TestSample2(t *testing.T) {
	N, m, x, y := 3, 2, 6, 7
	A := []int{1, 6, 9}
	expect := 36
	runSample(t, N, m, x, y, A, expect)
}
