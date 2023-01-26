package main

import "testing"

func runSample(t *testing.T, A []int, B []int, expect int) {
	res := solve(A, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{0, 3, 3}
	B := []int{0, 2, 2}
	expect := 1
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 3}
	B := []int{0, 1}
	expect := 3
	runSample(t, A, B, expect)
}
