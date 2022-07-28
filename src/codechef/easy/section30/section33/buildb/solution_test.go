package main

import "testing"

func runSample(t *testing.T, A, B []int, r int, expect int) {
	res := int(solve(A, B, r))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{10}
	B := []int{10}
	r := 2
	expect := 10
	runSample(t, A, B, r, expect)
}

func TestSample2(t *testing.T) {
	A := []int{10, 11}
	B := []int{10, 10}
	r := 2
	expect := 18
	runSample(t, A, B, r, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 2, 3}
	B := []int{1, 2, 3}
	r := 1
	expect := 4
	runSample(t, A, B, r, expect)
}
