package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 1, 2, 2, 3, 3, 3}
	expect := 6
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 1, 1, 5, 4, 1, 3, 1, 2, 2}
	expect := 7
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 2, 1, 3, 1, 4, 1, 5, 1}
	expect := 3
	runSample(t, A, expect)
}
