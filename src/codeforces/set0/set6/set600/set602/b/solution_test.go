package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3, 3, 2}
	expect := 4
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{5, 4, 5, 5, 6, 7, 8, 8, 8, 7, 6}
	expect := 5
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 1, 1, 1, 1}
	expect := 5
	runSample(t, A, expect)
}
