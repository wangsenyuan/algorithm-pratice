package main

import "testing"

func runSample(t *testing.T, P []int, expect bool) {
	res := solve(P)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	P := []int{3, 1, 2, 4}
	expect := true
	runSample(t, P, expect)
}

func TestSample2(t *testing.T) {
	P := []int{2, 3, 1}
	expect := false
	runSample(t, P, expect)
}

func TestSample3(t *testing.T) {
	P := []int{5, 2, 3, 6, 1, 4}
	expect := true
	runSample(t, P, expect)
}
