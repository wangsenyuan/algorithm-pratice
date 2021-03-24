package main

import "testing"

func runSample(t *testing.T, n int, k int, A []int, expect int) {
	res := solve(n, k, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 5, 3
	A := []int{1, 2, 3, 2, 1}
	expect := 2
	runSample(t, n, k, A, expect)
}

func TestSample2(t *testing.T) {
	n, k := 4, 2
	A := []int{1, 2, 3, 4}
	expect := 3
	runSample(t, n, k, A, expect)
}

func TestSample3(t *testing.T) {
	n, k := 10, 7
	A := []int{9, 6, 2, 4, 4, 8, 4, 6, 2, 1}
	expect := solve1(n, k, A)
	runSample(t, n, k, A, expect)
}
