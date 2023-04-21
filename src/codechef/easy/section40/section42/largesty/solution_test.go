package main

import "testing"

func runSample(t *testing.T, A []int, X int, expect int) {
	res := solve(A, X)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 1, 2}
	X := 5
	expect := 5
	runSample(t, A, X, expect)
}

func TestSample2(t *testing.T) {
	A := []int{4, 5, 2, 3, 5}
	X := 10
	expect := 10
	runSample(t, A, X, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 1, 3}
	X := 3
	expect := 1
	runSample(t, A, X, expect)
}
