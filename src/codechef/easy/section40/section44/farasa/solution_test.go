package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{6, 1, 7, 3, 7, 1, 8, 5, 7}
	expect := 28
	runSample(t, A, expect)
}
