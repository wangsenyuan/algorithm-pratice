package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 3, 2, 1, 4, 2}
	expect := 2
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{100, 100, 4, 100}
	expect := 1
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 2, 3, 3, 3, 2, 6, 6}
	expect := 2
	runSample(t, A, expect)
}
