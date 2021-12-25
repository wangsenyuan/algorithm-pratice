package main

import "testing"

func runSample(t *testing.T, A []int, expect bool) {
	res := solve(len(A), A)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3, 3}
	expect := true

	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 1, 3, 3}
	expect := false

	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 2 ,2 ,1 ,5}
	expect := false

	runSample(t, A, expect)
}
