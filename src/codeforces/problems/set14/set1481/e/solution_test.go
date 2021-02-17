package main

import "testing"

func runSample(t *testing.T, C []int, expect int) {
	res := solve(len(C), C)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	C := []int{1, 2, 2, 1, 3}
	expect := 2
	runSample(t, C, expect)
}

func TestSample2(t *testing.T) {
	C := []int{1, 2, 2, 1, 1}
	expect := 1
	runSample(t, C, expect)
}
