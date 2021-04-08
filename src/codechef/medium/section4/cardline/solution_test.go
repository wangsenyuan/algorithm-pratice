package main

import "testing"

func runSample(t *testing.T, A, B []int, expect int) {
	res := solve(A, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 8, 4, 2, 6, 1, 5, 7}
	B := []int{5, 2, 4, 3, 8, 7, 6, 1}
	expect := 4
	runSample(t, A, B, expect)
}
