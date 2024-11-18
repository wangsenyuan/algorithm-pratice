package main

import "testing"

func runSample(t *testing.T, a []int, expect bool) {
	res := solve(a)
	if res != expect {
		t.Errorf("Sample %v, expect %t, but got %t", a, expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 1, 1, 2, 5}
	expect := true
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{4, 5, 3}
	expect := true
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{10, 10}
	expect := true
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{1, 2, 3}
	expect := false
	runSample(t, a, expect)
}
