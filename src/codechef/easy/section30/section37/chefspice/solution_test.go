package main

import "testing"

func runSample(t *testing.T, n int, m int, k int, expect bool) {
	res := solve(n, m, k)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3627, 17424, 27008005, false)
}
func TestSample2(t *testing.T) {
	runSample(t, 3, 5, 13, true)
}
