package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2}
	expect := 2
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{5, 2, 4, 3, 1}
	expect := 3
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{3, 8, 8, 2, 9, 1, 6, 2, 8, 3}
	expect := 6
	runSample(t, A, expect)
}
