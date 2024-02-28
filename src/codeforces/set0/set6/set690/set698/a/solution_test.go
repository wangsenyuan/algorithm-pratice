package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 3, 2, 0}
	expect := 2
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 3, 3, 2, 1, 2, 3}
	expect := 0
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{2, 2}
	expect := 1
	runSample(t, A, expect)
}
