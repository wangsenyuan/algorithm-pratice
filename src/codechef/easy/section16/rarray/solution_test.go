package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(len(A), A)

	if res != int64(expect) {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 3, 2, 4}
	expect := 6
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 3}
	expect := 6
	runSample(t, A, expect)
}
