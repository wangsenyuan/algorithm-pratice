package main

import "testing"

func runSample(t *testing.T, a int, b int, expect int) {
	res := solve(a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 3, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 30, 20, 4)
}