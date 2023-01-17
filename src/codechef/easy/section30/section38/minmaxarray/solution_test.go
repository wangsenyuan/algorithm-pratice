package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 5, 7, 6}
	expect := 5
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{10, 3, 5, 7}
	expect := 10
	runSample(t, A, expect)
}
