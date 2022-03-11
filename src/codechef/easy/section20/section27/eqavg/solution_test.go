package main

import "testing"

func runSample(t *testing.T, k int, A []int, expect bool) {
	res := solve(k, A)

	if expect != (len(res) > 0) {
		t.Errorf("Sample expect %t, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 3
	A := []int{1, 1, 2, 2, 3}
	expect := true
	runSample(t, k, A, expect)
}

func TestSample2(t *testing.T) {
	k := 2
	A := []int{1, 1, 1, 1, 1, 2}
	expect := false
	runSample(t, k, A, expect)
}
