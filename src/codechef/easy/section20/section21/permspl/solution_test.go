package main

import "testing"

func runSample(t *testing.T, n int, P []int, expect bool) {
	res := solve(n, P)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 1
	P := []int{1}
	expect := true
	runSample(t, n, P, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	P := []int{3, 2, 1}
	expect := false
	runSample(t, n, P, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	P := []int{4, 3, 2, 1}
	expect := true
	runSample(t, n, P, expect)
}

func TestSample4(t *testing.T) {
	n := 5
	P := []int{1, 4, 3, 2, 5}
	expect := false
	runSample(t, n, P, expect)
}
