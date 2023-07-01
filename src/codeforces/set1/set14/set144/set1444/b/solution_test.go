package main

import "testing"

func runSample(t *testing.T, X []int, expect int) {
	res := solve(X)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	X := []int{2, 1, 2, 1}
	expect := 12
	runSample(t, X, expect)
}

func TestSample2(t *testing.T) {
	X := []int{13, 8, 35, 94, 9284, 34, 54, 69, 123, 846}
	expect := 2588544
	runSample(t, X, expect)
}
