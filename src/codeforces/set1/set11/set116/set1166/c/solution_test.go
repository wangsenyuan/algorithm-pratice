package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 5, -3}
	expect := 2
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{3, 6}
	expect := 1
	runSample(t, a, expect)
}
