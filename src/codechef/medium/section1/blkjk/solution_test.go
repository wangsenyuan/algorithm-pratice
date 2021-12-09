package main

import "testing"

func runSample(t *testing.T, A []int, X, Y int, expect int) {
	res := solve(A, X, Y)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	X, Y := 4, 5
	A := []int{1, 2, 3}
	expect := 1
	runSample(t, A, X, Y, expect)
}

func TestSample2(t *testing.T) {
	X, Y := 3, 4
	A := []int{1, 2, 3}
	expect := 0
	runSample(t, A, X, Y, expect)
}

func TestSample3(t *testing.T) {
	X, Y := 20, 30
	A := []int{40, 10}
	expect := -1
	runSample(t, A, X, Y, expect)
}
