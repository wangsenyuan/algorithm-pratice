package main

import "testing"

func runSample(t *testing.T, A []int, expect bool) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 4, 4}
	expect := true
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 3, 4}
	expect := true
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 2}
	expect := false
	runSample(t, A, expect)
}
