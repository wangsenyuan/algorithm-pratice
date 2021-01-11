package main

import "testing"

func runSample(t *testing.T, A, B []int, expect int) {
	res := solve(A, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 2}
	B := []int{5, 5, 5}
	expect := 2
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 3, 2, 4}
	B := []int{6, 7, 8}
	expect := 1
	runSample(t, A, B, expect)
}
