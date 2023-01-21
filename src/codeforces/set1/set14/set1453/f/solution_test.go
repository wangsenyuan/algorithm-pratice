package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	A := []int{1, 1, 1, 0}
	expect := 0
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	A := []int{4, 3, 2, 1, 0}
	expect := 3
	runSample(t, n, A, expect)
}

func TestSample3(t *testing.T) {
	n := 9
	A := []int{4, 1, 4, 2, 1, 0, 2, 1, 0}
	expect := 2
	runSample(t, n, A, expect)
}
