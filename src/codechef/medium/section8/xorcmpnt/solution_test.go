package main

import "testing"

func runSample(t *testing.T, k int, A []int, expect int) {
	res := solve(k, A)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", k, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 2
	A := []int{0}
	expect := 4
	runSample(t, k, A, expect)
}

func TestSample2(t *testing.T) {
	k := 2
	A := []int{2}
	expect := 2
	runSample(t, k, A, expect)
}
