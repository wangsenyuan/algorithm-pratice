package main

import "testing"

func runSample(t *testing.T, A, B, C []int, expect bool) {
	res := solve(A, B, C)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 4}
	B := []int{-4, 3}
	C := []int{1, 7}
	expect := true
	runSample(t, A, B, C, expect)
}
