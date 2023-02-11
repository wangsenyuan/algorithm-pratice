package main

import "testing"

func runSample(t *testing.T, k int, A []int, expect int) {
	res := solve(k, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 1
	A := []int{3, 0, 1, 2}
	expect := 0
	runSample(t, k, A, expect)
}

func TestSample2(t *testing.T) {
	k := 1
	A := []int{0, 2, 4, 5}
	expect := 1
	runSample(t, k, A, expect)
}
