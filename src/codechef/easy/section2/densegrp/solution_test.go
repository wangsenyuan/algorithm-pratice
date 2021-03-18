package main

import "testing"

func runSample(t *testing.T, x, y int, A, B, C, D []int, expect int) {
	res := solve(x, y, A, B, C, D)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 5}
	B := []int{2, 5}
	C := []int{4, 2}
	D := []int{5, 3}
	x, y := 1, 2
	expect := 2
	runSample(t, x, y, A, B, C, D, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1}
	B := []int{6}
	C := []int{3}
	D := []int{3}
	x, y := 3, 4
	expect := -1
	runSample(t, x, y, A, B, C, D, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1}
	B := []int{1}
	C := []int{1}
	D := []int{1}
	x, y := 2, 2
	expect := 0
	runSample(t, x, y, A, B, C, D, expect)
}
