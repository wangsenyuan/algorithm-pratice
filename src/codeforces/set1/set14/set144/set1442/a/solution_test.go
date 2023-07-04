package main

import "testing"

func runSample(t *testing.T, A []int, expect bool) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 1}
	expect := true
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{11, 7, 9, 6, 8}
	expect := true
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 3, 1, 3, 1}
	expect := false
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{5, 2, 1, 10}
	expect := true
	runSample(t, A, expect)
}
