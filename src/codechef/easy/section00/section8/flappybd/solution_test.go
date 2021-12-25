package main

import "testing"

func runSample(t *testing.T, h, n int, X []int, H []int, expect int) {
	res := solve(h, n, X, H)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, h := 1, 0
	X := []int{2}
	H := []int{1}
	expect := 2
	runSample(t, h, n, X, H, expect)
}

func TestSample2(t *testing.T) {
	n, h := 2, 1
	X := []int{1, 3}
	H := []int{1, 1}
	expect := 2
	runSample(t, h, n, X, H, expect)
}

func TestSample3(t *testing.T) {
	n, h := 5, 10
	X := []int{1, 2, 3, 4, 5}
	H := []int{10, 11, 12, 13, 15}
	expect := -1
	runSample(t, h, n, X, H, expect)
}
