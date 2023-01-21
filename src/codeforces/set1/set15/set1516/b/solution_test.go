package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect bool) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{962815211, 962815211, 962815211}
	var expect bool = true
	runSample(t, n, A, expect)
}
