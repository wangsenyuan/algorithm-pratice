package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{1, 2, 3}
	expect := 2
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	A := []int{2, 1, 1, 2}
	expect := 1
	runSample(t, n, A, expect)
}
