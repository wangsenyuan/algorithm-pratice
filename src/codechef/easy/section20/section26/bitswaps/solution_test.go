package main

import "testing"

func runSample(t *testing.T, A []int, expect bool) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{6, 4, 2}
	expect := true
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{9, 34, 4, 24, 1, 6}
	expect := true
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{9, 34, 24, 4, 1, 6}
	expect := false
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{1, 0}
	expect := false
	runSample(t, A, expect)
}
