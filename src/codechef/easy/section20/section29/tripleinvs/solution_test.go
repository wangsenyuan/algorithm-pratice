package main

import "testing"

func runSample(t *testing.T, A []int, expect bool) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{0, 1, 3, 2}
	expect := true
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{0, 1, 2, 3}
	expect := false
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{0, 0, 0, 0}
	expect := true
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{0, 1, 2, 1, 0}
	expect := false
	runSample(t, A, expect)
}
