package main

import "testing"

func runSample(t *testing.T, l int, r int, expect int) {
	res := solve(l, r)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 9, 8)
}

func TestSample2(t *testing.T) {
	runSample(t, 9, 10, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 1, 1000000000, 1111111110)
}
