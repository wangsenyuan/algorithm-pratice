package main

import "testing"

func runSample(t *testing.T, A []int, expect bool) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, -1, -1, 0}
	expect := true
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, -4, 3, 0}
	expect := false
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, -1, 1, -1}
	expect := false
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{1, 2, 3, 4, -10}
	expect := true
	runSample(t, A, expect)
}

func TestSample5(t *testing.T) {
	A := []int{2, -1, 1, -2, 0, 0, 0}
	expect := true
	runSample(t, A, expect)
}
