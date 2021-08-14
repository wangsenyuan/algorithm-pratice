package main

import "testing"

func runSample(t *testing.T, h int, k int, X []int, expect int) {
	res := solve(h, k, X)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	h, k := 2, 10
	X := []int{70, 78, 67, 75}
	expect := 2

	runSample(t, h, k, X, expect)
}

func TestSample2(t *testing.T) {
	h, k := 3, 10
	X := []int{18, 2, 19, 21, 33, 40, 26, 25}
	expect := 4

	runSample(t, h, k, X, expect)
}

func TestSample3(t *testing.T) {
	h, k := 2, 7
	X := []int{4, 9, 8, 16}
	expect := -1

	runSample(t, h, k, X, expect)
}
