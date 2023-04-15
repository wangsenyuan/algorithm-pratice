package main

import "testing"

func runSample(t *testing.T, A []int, B []int, expect int) {
	res := solve(A, B)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 0, 1}
	B := []int{1, 2, 2}
	expect := 2
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 2, 1, 0, 0, 1, 1, 1}
	B := []int{4, 4, 4, 7, 3, 2, 2, 1}
	expect := 3
	runSample(t, A, B, expect)
}
