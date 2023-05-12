package main

import "testing"

func runSample(t *testing.T, X []int, Y []int, expect int) {
	res := solve(X, Y)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	X := []int{1, 2, 4}
	Y := []int{3, 5, 6}
	expect := 2

	runSample(t, X, Y, expect)
}

func TestSample2(t *testing.T) {
	X := []int{1, 2, 5}
	Y := []int{3, 4, 6}
	expect := 1

	runSample(t, X, Y, expect)
}

func TestSample3(t *testing.T) {
	X := []int{1, 2, 3, 4}
	Y := []int{6, 7, 5, 8}
	expect := 3

	runSample(t, X, Y, expect)
}

func TestSample4(t *testing.T) {
	X := []int{1, 3, 5}
	Y := []int{2, 4, 6}
	expect := 0

	runSample(t, X, Y, expect)
}
