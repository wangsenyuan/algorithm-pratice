package main

import "testing"

func runSample(t *testing.T, a int, b int, expect int) {
	res := solve(a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 15, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, 13, 37, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, 8, 9, -1)
}

func TestSample4(t *testing.T) {
	runSample(t, 10009, 20000, 79)
}
