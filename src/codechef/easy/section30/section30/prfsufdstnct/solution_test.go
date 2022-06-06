package main

import "testing"

func runSample(t *testing.T, P, S []int, expect bool) {
	res := solve(P, S)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	P := []int{1, 2, 3}
	S := []int{3, 2, 1}
	expect := true
	runSample(t, P, S, expect)
}

func TestSample2(t *testing.T) {
	P := []int{2, 2, 3, 3}
	S := []int{3, 2, 2, 1}
	expect := false
	runSample(t, P, S, expect)
}

func TestSample3(t *testing.T) {
	P := []int{1, 1}
	S := []int{1, 1}
	expect := true
	runSample(t, P, S, expect)
}
