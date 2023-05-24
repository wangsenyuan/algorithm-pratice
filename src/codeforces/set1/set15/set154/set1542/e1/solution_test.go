package main

import "testing"

func runSample(t *testing.T, n int, mod int, expect int) {
	res := solve(n, mod)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 403458273, 17)
}
