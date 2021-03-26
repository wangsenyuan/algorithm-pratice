package main

import "testing"

func runSample(t *testing.T, k int, P, Q []int, expect int) {
	res := solve(k, P, Q)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 2
	P := []int{2, 4, 3, 1}
	Q := []int{1, 2, 4, 3}
	expect := 2
	runSample(t, k, P, Q, expect)
}

func TestSample2(t *testing.T) {
	k := 3
	P := []int{1, 2, 3}
	Q := []int{1, 3, 2}
	expect := -1
	runSample(t, k, P, Q, expect)
}
