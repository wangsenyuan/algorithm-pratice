package main

import "testing"

func runSample(t *testing.T, n int, k int, expect int) {
	res := solve(n, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 2, 9)
}

func TestSample2(t *testing.T) {
	runSample(t, 100, 1, 100)
}

func TestSample3(t *testing.T) {
	runSample(t, 1000000000000, 1377, 999244007)
}

func TestSample4(t *testing.T) {
	runSample(t, 193216427128336546, 31250, 635886063)
}

func TestSample5(t *testing.T) {
	runSample(t, 284291990143008233, 31250, 202309185)
}
