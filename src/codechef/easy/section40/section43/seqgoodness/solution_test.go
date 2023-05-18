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
	expect := 3
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{3, 2, 5, 5, 3}
	expect := 5
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 1, 2, 1}
	expect := 17
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{6, 4, 5, 2, 3, 1}
	expect := 63
	runSample(t, A, expect)
}
