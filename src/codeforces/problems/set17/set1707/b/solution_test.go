package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 10, 100}
	expect := 81
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{4, 8, 9, 13}
	expect := 3
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{0, 0, 0, 8, 13}
	expect := 1
	runSample(t, A, expect)
}
