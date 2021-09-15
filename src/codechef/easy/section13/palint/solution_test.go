package main

import "testing"

func runSample(t *testing.T, A []int, X int, expect int, flips int) {
	res, cnt := solve(A, X)

	if res != expect || cnt != flips {
		t.Errorf("Sample expect %d %d, but got %d %d", expect, flips, res, cnt)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3}
	X := 2
	expect := 2
	flips := 1
	runSample(t, A, X, expect, flips)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 3, 4, 5}
	X := 100
	expect := 1
	flips := 0
	runSample(t, A, X, expect, flips)
}

func TestSample3(t *testing.T) {
	A := []int{2, 2, 6, 6}
	X := 1
	expect := 2
	flips := 0
	runSample(t, A, X, expect, flips)
}
