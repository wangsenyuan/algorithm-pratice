package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := int(solve(len(A), A))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 1}
	expect := 4
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 4}
	expect := 7
	runSample(t, A, expect)
}
