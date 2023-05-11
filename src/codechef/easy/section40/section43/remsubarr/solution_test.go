package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 1, 3}
	expect := 1
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 3, 4, 5, 6, 7}
	expect := 6
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{3, 1, 4, 6, 5, 2}
	expect := 3
	runSample(t, A, expect)
}
