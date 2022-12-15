package main

import "testing"

func runSample(t *testing.T, P []int, expect int) {
	res := solve(P)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	P := []int{1, 2, 3}
	expect := 3
	runSample(t, P, expect)
}

func TestSample2(t *testing.T) {
	P := []int{2, 3, 4, 1}
	expect := 5
	runSample(t, P, expect)
}

func TestSample3(t *testing.T) {
	P := []int{3, 4, 2, 1}
	expect := 6
	runSample(t, P, expect)
}

func TestSample4(t *testing.T) {
	P := []int{4, 5, 2, 1, 3}
	expect := 6
	runSample(t, P, expect)
}
