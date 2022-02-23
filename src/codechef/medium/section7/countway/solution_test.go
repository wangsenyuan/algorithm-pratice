package main

import "testing"

func runSample(t *testing.T, k int, x []int, expect int) {
	res := solve(x, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 2, 2, 3}
	k := 2
	expect := 3
	runSample(t, k, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 2, 2, 2, 2}
	k := 1
	expect := 1
	runSample(t, k, A, expect)
}
