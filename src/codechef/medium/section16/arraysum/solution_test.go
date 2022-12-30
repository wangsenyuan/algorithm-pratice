package main

import "testing"

func runSample(t *testing.T, A []int, B []int, expect int) {
	res := solve(A, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 5, 2, 4, 7}
	B := []int{7, 9, 6}
	expect := 4
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{7, 7, 7, 7}
	B := []int{3, 4}
	expect := 6
	runSample(t, A, B, expect)
}
