package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 1, 2}
	expect := 2
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{3, 3, 3, 3}
	expect := 1
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{5, 5, 10, 20}
	expect := 4
	runSample(t, A, expect)
}
