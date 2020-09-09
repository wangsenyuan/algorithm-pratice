package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{1, 2, 3, 4, 5}
	expect := 0
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 7
	A := []int{3, 2, 4, 5, 1, 6, 7}
	expect := 2
	runSample(t, n, A, expect)
}
