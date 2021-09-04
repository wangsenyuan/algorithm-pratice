package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{9, 15, 8, 13, 8}
	expect := -1
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{20, 16, 13, 9, 17, 11, 15, 8, 7}
	expect := 4
	runSample(t, A, expect)
}
