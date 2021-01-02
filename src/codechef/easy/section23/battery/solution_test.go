package main

import "testing"

func runSample(t *testing.T, X int, A []int, expect bool) {
	res := solve(X, A)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	X := 8
	A := []int{5, 1, 4, 2}
	expect := true
	runSample(t, X, A, expect)
}

func TestSample2(t *testing.T) {
	X := 4
	A := []int{3, 1, 2}
	expect := true
	runSample(t, X, A, expect)
}

func TestSample3(t *testing.T) {
	X := 7
	A := []int{5, 1, 5}
	expect := false
	runSample(t, X, A, expect)
}
