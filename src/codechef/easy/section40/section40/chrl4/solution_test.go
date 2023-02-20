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
	A := []int{1, 2, 3, 4}
	expect := 8
	runSample(t, A, k, expect)
}
