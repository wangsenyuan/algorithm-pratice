package main

import "testing"

func runSample(t *testing.T, n int, k int, A []int, expect int) {
	res := solve(n, k, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 4, 1
	A := []int{1, 1, 2, 3}
	expect := 3
	runSample(t, n, k, A, expect)
}

func TestSample2(t *testing.T) {
	n, k := 9, 2
	A := []int{1, 2, 3, 2, 4, 5, 6, 7, 5}
	expect := 5
	runSample(t, n, k, A, expect)
}

func TestSample3(t *testing.T) {
	n, k := 5, 5
	A := []int{1, 1, 1, 1, 1}
	expect := 5
	runSample(t, n, k, A, expect)
}

func TestSample4(t *testing.T) {
	n, k := 10, 1
	A := []int{1, 2, 1, 2, 1, 2, 1, 2, 1, 2}
	expect := 6
	runSample(t, n, k, A, expect)
}
