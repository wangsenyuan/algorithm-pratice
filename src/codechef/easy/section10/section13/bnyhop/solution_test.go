package main

import "testing"

func runSample(t *testing.T, A []int, P []int, expect int) {
	res := solve(len(A), A, P)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{10, 15, 5}
	P := []int{1, 1}
	expect := 3
	runSample(t, A, P, expect)
}

func TestSample2(t *testing.T) {
	A := []int{50, 10, 20, 40, 30}
	P := []int{1, 1, 3, 3}
	expect := 8
	runSample(t, A, P, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	P := []int{1, 1, 2, 2, 2, 2, 7, 7}
	expect := 28
	runSample(t, A, P, expect)
}
