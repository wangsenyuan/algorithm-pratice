package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)
	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 1)
}