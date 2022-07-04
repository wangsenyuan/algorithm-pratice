package main

import "testing"

func runSample(t *testing.T, u int, v int, expect int) {
	res := int(solve(u, v))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 6, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 3, 5)
}