package main

import "testing"

func runSample(t *testing.T, a, b, d int, expect int) {
	res := solve(a, b, d)
	if expect < 0 {
		if res < 0 {
			return
		}
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}

	a |= res
	b |= res
	if a%d != 0 || b%d != 0 {
		t.Fatalf("Sample result %d, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 12, 39, 5, 18)
}

func TestSample2(t *testing.T) {
	runSample(t, 987654321, 123456789, 999999999, 184470016815529983)
}
