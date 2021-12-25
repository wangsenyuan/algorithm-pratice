package main

import "testing"

func runSample(t *testing.T, k int, A []int, expect bool) {
	res := solve(len(A), k, A)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 4, 3}
	k := 2
	expect := true
	runSample(t, k, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{6, 6, 8, 7, 3}
	k := 6
	expect := false
	runSample(t, k, A, expect)
}
