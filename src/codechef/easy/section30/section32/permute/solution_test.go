package main

import "testing"

func runSample(t *testing.T, n int, m int, expect int) {
	res := solve(n, m)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 5, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 8, 72)
}