package main

import "testing"

func runSample(t *testing.T, A []int, expect bool) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{6, 11}
	expect := true
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 18, 22}
	expect := false
	runSample(t, A, expect)
}
