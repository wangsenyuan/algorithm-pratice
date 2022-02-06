package main

import "testing"

func runSample(t *testing.T, A, B []int, X, Y int, expect int) {
	res := int(solve(A, B, X, Y))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3, 4, 5}
	B := []int{5, 4, 3, 2, 1}
	X, Y := 3, 3
	expect := 21
	runSample(t, A, B, X, Y, expect)
}
