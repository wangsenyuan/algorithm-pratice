package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2}
	expect := 6
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 3}
	expect := 20
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{5, 4, 1, 3}
	expect := 82
	runSample(t, A, expect)
}
