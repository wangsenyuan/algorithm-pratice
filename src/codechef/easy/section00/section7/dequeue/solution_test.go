package main

import "testing"

func runSample(t *testing.T, n int, m int, Q []int, expect int) {
	res := solve(n, m, Q)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 2, 2
	Q := []int{1, 2}
	expect := 2
	runSample(t, n, m, Q, expect)
}

func TestSample2(t *testing.T) {
	n, m := 3, 6
	Q := []int{1, 1, 1, 3, 2, 2}
	expect := 4
	runSample(t, n, m, Q, expect)
}
