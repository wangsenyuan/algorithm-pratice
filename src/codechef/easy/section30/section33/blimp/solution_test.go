package main

import "testing"

func runSample(t *testing.T, X int, Y int, A []int, expect int) {
	res := solve(X, Y, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 5, 4, 4}
	X, Y := 4, 4
	expect := 2
	runSample(t, X, Y, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 4, 3, 3, 5}
	X, Y := 4, 3
	expect := 2
	runSample(t, X, Y, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{3, 1, 3, 9}
	X, Y := 3, 1
	expect := 3
	runSample(t, X, Y, A, expect)
}
