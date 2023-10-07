package main

import "testing"

func runSample(t *testing.T, n int, m int, expect int) {
	res := solve(n, m)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 3, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 2, 26)
}

func TestSample3(t *testing.T) {
	runSample(t, 4, 6, 1494)
}

func TestSample4(t *testing.T) {
	runSample(t, 1337, 424242424242, 119112628)
}
