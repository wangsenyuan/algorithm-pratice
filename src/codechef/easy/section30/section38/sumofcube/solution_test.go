package main

import "testing"

func runSample(t *testing.T, P []int, expect int) {
	res := solve(P)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	P := []int{1, 1}
	expect := 10
	runSample(t, P, expect)
}

func TestSample2(t *testing.T) {
	P := []int{1, 2, 1}
	expect := 128
	runSample(t, P, expect)
}

func TestSample3(t *testing.T) {
	P := []int{8, 5, 6, 2, 3}
	expect := 42621
	runSample(t, P, expect)
}
