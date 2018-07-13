package main

import "testing"

func runSample(t *testing.T, N, m, X, Y, Z int, A []int, expect int64) {
	res := solve(N, m, X, Y, Z, A)
	if res != expect {
		t.Errorf("Sample %d %d %d %d %d %v, expect %d, but got %d", N, m, X, Y, Z, A, expect, res)
	}
}

func TestSample2(t *testing.T) {
	N, m, X, Y, Z := 6, 2, 2, 1000000000, 6
	A := []int{1, 2}
	expect := int64(13)
	runSample(t, N, m, X, Y, Z, A, expect)
}

func TestSample1(t *testing.T) {
	N, m, X, Y, Z := 5, 5, 0, 0, 5
	A := []int{1, 2, 1, 2, 3}
	expect := int64(15)
	runSample(t, N, m, X, Y, Z, A, expect)
}
