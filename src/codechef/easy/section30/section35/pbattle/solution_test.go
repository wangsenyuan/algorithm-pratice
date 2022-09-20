package main

import "testing"

func runSample(t *testing.T, A []int, B []int, expect int) {
	res := solve(A, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 3, 5, 4, 1}
	B := []int{4, 2, 1, 5, 6}
	expect := 3
	runSample(t, A, B, expect)
}
