package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{4, 31, 25, 50, 30, 20, 34, 46, 42, 16, 15, 16}
	expect := 3
	runSample(t, A, expect)
}
