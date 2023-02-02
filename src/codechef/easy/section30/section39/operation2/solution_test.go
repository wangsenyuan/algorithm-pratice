package main

import "testing"

func runSample(t *testing.T, C []int, expect int) {
	res := solve(C)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	C := []int{2, 0, 1, 3}
	expect := 3
	runSample(t, C, expect)
}

func TestSample2(t *testing.T) {
	C := []int{1, 2, 5, 6}
	expect := 7
	runSample(t, C, expect)
}
