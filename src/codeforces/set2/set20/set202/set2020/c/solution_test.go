package main

import "testing"

func runSample(t *testing.T, b, c, d int, expect int) {
	a := solve(b, c, d)

	if a == expect {
		return
	}
	if expect < 0 || a < 0 || (a|b)-(a&c) != d {
		t.Fatalf("Sample expect %d, but got %d", expect, a)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 2, 2, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 2, 6, -1)
}

func TestSample4(t *testing.T) {
	runSample(t, 10, 2, 14, 12)
}
