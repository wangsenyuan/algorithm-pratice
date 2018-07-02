package main

import "testing"

func runSample(t *testing.T, n int, m int, k int, T []int, P []int, expect int) {
	res := solve(n, m, k, T, P)

	if res != expect {
		t.Errorf("Sample %d %d %d %v %v, expect %d, but got %d", n, m, k, T, P, expect, res)
	}
}

func TestSample1(t *testing.T) {
	m, n, k := 3, 2, 5
	T := []int{5, 1, 2}
	P := []int{1, 2, 1}
	expect := 4
	runSample(t, n, m, k, T, P, expect)
}
