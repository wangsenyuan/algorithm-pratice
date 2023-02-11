package main

import "testing"

func runSample(t *testing.T, P []int, expect int) {
	res := solve(P)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{0, 1, 3, 2}
	expect := 2
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{0, 1, 2, 3, 5, 6, 4}
	expect := 4
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{0, 3, 2, 1, 4}
	expect := 1
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{2, 0, 1}
	expect := 0
	runSample(t, A, expect)
}
