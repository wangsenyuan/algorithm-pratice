package main

import "testing"

func runSample(t *testing.T, n int, size int) {
	res := solve(n)

	if len(res) != size {
		t.Errorf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	n := 9
	size := 7
	runSample(t, n, size)
}

func TestSample2(t *testing.T) {
	n := 7
	size := 4
	runSample(t, n, size)
}
