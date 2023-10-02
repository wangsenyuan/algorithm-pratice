package main

import "testing"

func runSample(t *testing.T, n int, m int, expect int) {
	res := solve(n, m)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 6, 11, 60)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 43, 200)
}

func TestSample3(t *testing.T) {
	runSample(t, 13, 5, 65)
}

func TestSample4(t *testing.T) {
	runSample(t, 4, 16, 60)
}
