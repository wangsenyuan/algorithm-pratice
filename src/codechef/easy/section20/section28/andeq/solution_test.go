package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{0, 0, 0, 1}
	expect := 1
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 3, 4, 5, 6}
	expect := 4
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{2, 28, 3, 22}
	expect := 3
	runSample(t, A, expect)
}
