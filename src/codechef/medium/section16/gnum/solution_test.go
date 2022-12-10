package main

import "testing"

func runSample(t *testing.T, A []int, B []int, expect int) {
	res := solve(A, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 5, 6, 14}
	B := []int{3, 4, 7, 10}
	expect := 3
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 3}
	B := []int{5, 7}
	expect := 0
	runSample(t, A, B, expect)
}
