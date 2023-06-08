package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 1, 4}
	expect := 4
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 1, 3}
	expect := 2
	runSample(t, A, expect)
}
