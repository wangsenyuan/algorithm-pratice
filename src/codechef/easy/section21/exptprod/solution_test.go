package main

import "testing"

func runSample(t *testing.T, n int, k int, A []int, expect int) {
	res := solve(n, k, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 2
	A := []int{2, 1, 2}
	expect := 444444449
	runSample(t, n, k, A, expect)
}

func TestSample2(t *testing.T) {
	n, k := 2, 3
	A := []int{1, 1}
	expect := 1
	runSample(t, n, k, A, expect)
}
