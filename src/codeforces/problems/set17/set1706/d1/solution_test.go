package main

import "testing"

func runSample(t *testing.T, A []int, k int, expect int) {
	res := solve(A, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 2
	A := []int{4, 5, 6, 8, 11}
	expect := 2
	runSample(t, A, k, expect)
}
