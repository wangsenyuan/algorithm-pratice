package main

import "testing"

func runSample(t *testing.T, A []int, expect bool) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3}
	expect := true
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 3, 4, 2}
	expect := true
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{2, 3, 6, 2, 0, 2}
	expect := false
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{2, 4, 6, 4, 2}
	expect := true
	runSample(t, A, expect)
}
