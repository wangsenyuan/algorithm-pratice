package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 3, 7, 15}
	expect := 14
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 1, 2, 2}
	expect := 3
	runSample(t, A, expect)
}
