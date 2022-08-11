package main

import "testing"

func runSample(t *testing.T, n int, m int, expect int) {
	res := solve(int64(n), int64(m))
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 3, 32)
}
