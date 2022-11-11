package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 2, 5}
	expect := 3
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{10, 20, 30, 40}
	expect := 4
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{14, 11, 9, 3, 7}
	expect := 4
	runSample(t, A, expect)
}
