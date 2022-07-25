package main

import "testing"

func runSample(t *testing.T, A []int, k int, expect int) {
	res := int(solve(A, k))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3, 4}
	k := 2
	expect := 1
	runSample(t, A, k, expect)
}
