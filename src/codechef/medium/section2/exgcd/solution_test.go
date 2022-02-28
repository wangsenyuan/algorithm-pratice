package main

import "testing"

func runSample(t *testing.T, A, B []int, expect int) {
	res := solve(A, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 3}
	B := []int{4, 5}
	expect := 333333334
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 1, 1}
	B := []int{12, 12, 12}
	expect := 722222226
	runSample(t, A, B, expect)
}
