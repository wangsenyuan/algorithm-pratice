package main

import "testing"

func runSample(t *testing.T, A, B int, C []int, D []int, expect int) {
	res := solve(A, B, C, D)
	if res != expect {
		t.Errorf("Sample %d %d %v %v, expect %d, but got %d", A, B, C, D, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A, B := 1, 3
	C := []int{4, 2, 9}
	D := []int{5, 6, 7}
	expect := 1
	runSample(t, A, B, C, D, expect)
}

func TestSample2(t *testing.T) {
	A, B := 1, 3
	C := []int{4, 2, 9}
	D := []int{5, 10, 7}
	expect := 0
	runSample(t, A, B, C, D, expect)
}

func TestSample3(t *testing.T) {
	A, B := 3, 3
	C := []int{7, 14, 11, 4, 15, 5, 20, 1, 17}
	D := []int{2, 13, 16, 9, 19, 6, 12, 8, 10}
	expect := 3
	runSample(t, A, B, C, D, expect)
}
