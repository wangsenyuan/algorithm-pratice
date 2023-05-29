package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{48, 13, 22, -15, 16, 35}
	expect := 13
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{-1, 0, 1, -1, 0, 1, -1, 0}
	expect := 2
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{0, 1, 2, 99, 4, 99, 99, 7, 99, 99, 99, 11, 12, 99, 99, 15, 16, 99, 18,
		19, 99, 99, 22, 99, 99, 25, 99, 99, 28, 99, 30, 31, 32, 99, 34, 35, 99, 99}
	expect := -1
	runSample(t, A, expect)
}
