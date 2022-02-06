package main

import "testing"

func runSample(t *testing.T, W, H int, X []int, Y []int, expect int) {
	res := solve(W, H, X, Y)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	W, H := 10, 10
	A := []int{3, 6, 8}
	B := []int{1, 6, 10}
	expect := 3
	runSample(t, W, H, A, B, expect)
}
