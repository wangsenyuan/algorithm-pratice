package main

import "testing"

func runSample(t *testing.T, m int, n int, expect int) {
	res := solve(m, n)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 2, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 2, 12)
}
