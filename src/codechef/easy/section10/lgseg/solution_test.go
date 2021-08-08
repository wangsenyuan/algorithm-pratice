package main

import "testing"

func runSample(t *testing.T, n int, k int, s int, A []int, expect int) {
	res := solve(n, k, s, A)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k, s := 5, 2, 5
	A := []int{1, 3, 2, 1, 5}
	expect := 4
	runSample(t, n, k, s, A, expect)
}

func TestSample2(t *testing.T) {
	n, k, s := 5, 3, 5
	A := []int{5, 1, 5, 1, 1}
	expect := 4
	runSample(t, n, k, s, A, expect)
}
