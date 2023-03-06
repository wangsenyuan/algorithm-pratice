package main

import "testing"

func runSample(t *testing.T, n int, fmt bool, expect int) {
	ok, res := solve(n)

	if ok != fmt || res != expect {
		t.Errorf("Sample expect %t %d, but got %t %d", fmt, expect, ok, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 957, true, 7493)
}

func TestSample2(t *testing.T) {
	runSample(t, 10000, true, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, 3, false, 1)
}

func TestSample4(t *testing.T) {
	runSample(t, 4, false, 2)
}
