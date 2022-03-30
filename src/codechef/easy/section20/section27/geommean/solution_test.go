package main

import "testing"

func runSample(t *testing.T, X int, A []int, expect int) {
	res := int(solve(X, A))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	X := 3
	A := []int{3, 3, 3}
	expect := 6
	runSample(t, X, A, expect)
}

func TestSample2(t *testing.T) {
	X := 4
	A := []int{1, 2, 3, 4}
	expect := 1
	runSample(t, X, A, expect)
}

func TestSample3(t *testing.T) {
	X := 54
	A := []int{36, 81, 54, 54}
	expect := 6
	runSample(t, X, A, expect)
}
