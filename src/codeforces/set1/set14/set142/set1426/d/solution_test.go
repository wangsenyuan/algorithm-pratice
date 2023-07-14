package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{-1, 1, -1, 1, -1, 1, 1, -1, -1}
	expect := 6
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{16, -5, -11, -15, 10, 5, 4, -4}
	expect := 3
	runSample(t, A, expect)
}
