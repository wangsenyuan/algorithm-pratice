package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", a, expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	expect := 1
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{0, 2, 2, 2, 4, 6}
	expect := 4
	runSample(t, a, expect)
}
