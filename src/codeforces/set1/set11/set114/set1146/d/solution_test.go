package main

import "testing"

func runSample(t *testing.T, m, a int, b int, expect int) {
	res := solve(a, b, m)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 7, 5, 3, 19)
}

func TestSample2(t *testing.T) {
	runSample(t, 1000000000, 1, 2019, 500000001500000001)
}
