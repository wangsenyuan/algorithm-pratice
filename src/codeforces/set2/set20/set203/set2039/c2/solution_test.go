package main

import "testing"

func runSample(t *testing.T, x int, m int, expect int) {
	res := solve(x, m)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 7, 10, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 3, 2)
}