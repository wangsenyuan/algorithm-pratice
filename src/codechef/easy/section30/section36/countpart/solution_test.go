package main

import "testing"

func runSample(t *testing.T, P []int, expect int) {
	res := solve(P)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	P := []int{2, 1, 4, 3}
	expect := 6
	runSample(t, P, expect)
}

func TestSample2(t *testing.T) {
	P := []int{2, 1}
	expect := 2
	runSample(t, P, expect)
}

func TestSample3(t *testing.T) {
	P := []int{2, 1, 4}
	// 2, 1, 4
	// 2 1, 4
	// 2, 1 4
	expect := 3
	runSample(t, P, expect)
}

func TestSample4(t *testing.T) {
	P := []int{1, 2, 3, 4, 5}
	// 2, 1, 4
	// 2 1, 4
	// 2, 1 4
	expect := 16
	runSample(t, P, expect)
}
