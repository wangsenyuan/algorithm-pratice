package main

import "testing"

func runSample(t *testing.T, A, B []int, expect int) {
	res := solve(A, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{6, 4, 5}
	B := []int{1, 3}
	expect := 4
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 2, 4}
	B := []int{1, 3}
	expect := 5
	runSample(t, A, B, expect)
}
