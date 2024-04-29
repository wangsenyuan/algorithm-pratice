package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", a, expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{0, 1, 0, 3, 0, 2}
	expect := 6
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{0, 1, 2, 1, 2}
	expect := 1
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{0, 1, 1, 2, 2}
	expect := 0
	runSample(t, a, expect)
}
