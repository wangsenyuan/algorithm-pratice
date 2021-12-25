package main

import "testing"

func runSample(t *testing.T, A []int, expect bool) {
	res := solve(len(A), A)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 0, 3}
	expect := true
	runSample(t, A, expect)
}
