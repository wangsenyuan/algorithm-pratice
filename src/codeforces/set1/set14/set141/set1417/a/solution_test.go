package main

import "testing"

func runSample(t *testing.T, n int, k int, A []int, expect int) {
	res := solve(n, k, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 2, 2
	A := []int{1, 1}
	expect := 1
	runSample(t, n, k, A, expect)
}

func TestSample2(t *testing.T) {
	n, k := 3, 5
	A := []int{1, 2, 3}
	expect := 5
	runSample(t, n, k, A, expect)
}

func TestSample3(t *testing.T) {
	n, k := 3, 7
	A := []int{2, 2, 3}
	expect := 4
	runSample(t, n, k, A, expect)
}
