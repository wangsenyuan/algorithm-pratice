package main

import "testing"

func runSample(t *testing.T, C []int, expect int) {
	res := solve(C)

	if res != expect {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	C := []int{-1, 1, -2, 1}
	expect := 1
	runSample(t, C, expect)
}

func TestSample2(t *testing.T) {
	C := []int{0, -1, 1, 2, -5}
	expect := 2
	runSample(t, C, expect)
}
