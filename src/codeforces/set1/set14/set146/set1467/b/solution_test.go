package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 5, 3}
	expect := 0
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 2, 2, 2, 2}
	expect := 0
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 6, 2, 5, 2, 10}
	expect := 1
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{1, 6, 2, 5, 1}
	expect := 0
	runSample(t, A, expect)
}
