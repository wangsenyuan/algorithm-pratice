package main

import "testing"

func runSample(t *testing.T, n int, m int, A []int, expect int) {
	res := solve(n, m, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 4, 3
	A := []int{1, 2, 1}
	expect := 2
	runSample(t, n, m, A, expect)
}

func TestSample2(t *testing.T) {
	n, m := 4, 6
	A := []int{1, 2, 1}
	expect := 0
	runSample(t, n, m, A, expect)
}

func TestSample3(t *testing.T) {
	n, m := 3, 2
	A := []int{2, 2}
	expect := 0
	runSample(t, n, m, A, expect)
}
