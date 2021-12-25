package main

import "testing"

func runSample(t *testing.T, n int, A []int, X int, p int, k int, expect int) {
	res := solve(n, A, X, p, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, X, p, k := 5, 4, 3, 4
	A := []int{4, 9, 7, 0, 8}
	expect := 1
	runSample(t, n, A, X, p, k, expect)
}

func TestSample2(t *testing.T) {
	n, X, p, k := 2, 25, 1, 2
	A := []int{100, 20}
	expect := -1
	runSample(t, n, A, X, p, k, expect)
}
