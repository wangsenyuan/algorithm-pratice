package main

import "testing"

func runSample(t *testing.T, n int, H []int, IQ []int, expect int) {
	res := solve(n, H, IQ)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	H := []int{1, 2, 3}
	IQ := []int{3, 2, 1}
	expect := 3
	runSample(t, n, H, IQ, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	H := []int{1, 3, 2, 4}
	IQ := []int{5, 6, 4, 4}
	expect := 2
	runSample(t, n, H, IQ, expect)
}
