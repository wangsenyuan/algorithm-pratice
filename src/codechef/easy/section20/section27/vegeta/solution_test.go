package main

import "testing"

func runSample(t *testing.T, n int, m int, expect int64) {
	res := solve(n, m)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 7, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 1023, 2177)
}
