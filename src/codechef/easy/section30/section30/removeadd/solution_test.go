package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 3, 3, 5, 1, 2, 4, 1}
	expect := 4
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 3, 4}
	expect := 0
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{10, 20, 10, 20, 30}
	expect := 2
	runSample(t, A, expect)
}
