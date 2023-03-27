package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3}
	expect := 5
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 1, 1, 1}
	expect := 15
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 3, 5, 15, 3, 105, 35}
	expect := 100
	runSample(t, A, expect)
}
