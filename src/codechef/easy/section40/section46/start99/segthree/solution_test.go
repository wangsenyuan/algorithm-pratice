package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3}
	expect := 0
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 3, 10, 25, 12, 7, 10, 12, 1, 46}
	expect := 3
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{10, 12, 15, 16, 17, 200, 132}
	expect := 4
	runSample(t, A, expect)
}
